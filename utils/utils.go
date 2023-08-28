package utils

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var Frontend string = "templates/layouts/frontend.html"
var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func CrearMensajesFlash(response http.ResponseWriter, request *http.Request, css string, message string) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, err := Store.Get(request, "flash-session")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	// Set some session values.
	session.AddFlash(css, "css")
	session.AddFlash(message, "message")

	// Save it before we write to the response/return from the handler.
	session.Save(request, response)
}

func RetornarMensajesFlash(response http.ResponseWriter, request *http.Request) (string, string) {
	// Get a session. Get() always returns a session, even if empty.
	session, _ := Store.Get(request, "flash-session")
	// Set some session values.
	css := session.Flashes("css")
	session.Save(request, response)
	css_session := ""
	if len(css) == 0 {
		css_session = ""
	} else {
		css_session = css[0].(string)
	}
	msg := session.Flashes("message")
	session.Save(request, response)
	msg_session := ""
	if len(msg) == 0 {
		msg_session = ""
	} else {
		msg_session = msg[0].(string)
	}
	return css_session, msg_session
}
