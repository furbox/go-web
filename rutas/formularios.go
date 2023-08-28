package rutas

import (
	"fmt"
	"go_web/utils"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Forms(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("views/forms/forms.html", utils.Frontend))
	css_session, msg_session := utils.RetornarMensajesFlash(response, request)
	data := map[string]string{
		"css":     css_session,
		"message": msg_session,
	}
	template.Execute(response, data)
}
func FormPost(response http.ResponseWriter, request *http.Request) {
	message := ""
	if utils.Regex_correo.FindStringSubmatch(request.FormValue("email")) == nil {
		message = message + "El campo email no es valido"
	}
	if !utils.ValidaPassword(request.FormValue("pass")) {
		message = message + "La contraseña debe tener al menos 8 caracteres, una mayuscula, una minuscula y un numero"
	}
	if len(request.FormValue("name")) == 0 {
		message = message + "El campo nombre es obligatorio"
	}
	if len(request.FormValue("phone")) != 10 {
		message = message + "El campo phone debe tener al menos 10 caracteres"
	}

	if message != "" {
		utils.CrearMensajesFlash(response, request, "alert-danger", message)
		http.Redirect(response, request, "/forms", http.StatusSeeOther)
	}

	/*
		template := template.Must(template.ParseFiles("views/forms/form-post.html", utils.Frontend))
		template.Execute(response, nil)
	*/
}
func Upload(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("views/forms/uploads.html", utils.Frontend))
	css_session, msg_session := utils.RetornarMensajesFlash(response, request)
	data := map[string]string{
		"css":     css_session,
		"message": msg_session,
	}
	template.Execute(response, data)
}
func UploadPost(response http.ResponseWriter, request *http.Request) {
	file, handler, err := request.FormFile("fileup")
	if err != nil {
		utils.CrearMensajesFlash(response, request, "alert-danger", "Error al subir el archivo")
		http.Redirect(response, request, "/uploads", http.StatusSeeOther)
	}
	fmt.Println(handler.Filename)
	var extensionWithDot = filepath.Ext(handler.Filename)
	extension := strings.TrimPrefix(extensionWithDot, ".")
	time := strings.Split(time.Now().String(), " ")
	fileName := string(time[4][6:14]) + "." + extension
	var path string = "public/uploads/" + fileName
	//guardar archivo
	f, errSaveFile := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
	if errSaveFile != nil {
		utils.CrearMensajesFlash(response, request, "alert-danger", "Error al subir el archivo")
		http.Redirect(response, request, "/uploads", http.StatusSeeOther)
	}
	_, errCopy := io.Copy(f, file)
	if errCopy != nil {
		utils.CrearMensajesFlash(response, request, "alert-danger", "Error al subir el archivo")
		http.Redirect(response, request, "/uploads", http.StatusSeeOther)
	}
	//guardar en db
	utils.CrearMensajesFlash(response, request, "alert-success", "Archivo subido correctamente")
	http.Redirect(response, request, "/uploads", http.StatusSeeOther)
}
