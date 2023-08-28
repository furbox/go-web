package rutas

import (
	"go_web/utils"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func Page404(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("views/error404.html", utils.Frontend))
	template.Execute(response, nil)
}
func Home(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("views/home.html", utils.Frontend))
	template.Execute(response, nil)
}
func Nosotros(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("views/nosotros.html", utils.Frontend))
	template.Execute(response, nil)
}
func Params(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("views/params.html", utils.Frontend))
	vars := mux.Vars(request)
	data := map[string]string{
		"id":   vars["id"],
		"slug": vars["slug"],
	}
	template.Execute(response, data)
}
func Queries(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("views/queries.html", utils.Frontend))
	query := request.URL.Query()
	data := map[string]string{
		"id":   query.Get("id"),
		"name": query.Get("name"),
	}
	template.Execute(response, data)
}

type Habilidad struct {
	Nombre string
}
type Datos struct {
	Nombre      string
	Edad        int
	Perfil      int
	Habilidades []Habilidad
}

func Estructuras(response http.ResponseWriter, request *http.Request) {
	habilidad1 := Habilidad{"Inteligencia"}
	habilidad2 := Habilidad{"Fuerza"}
	habilidad3 := Habilidad{"Velocidad"}
	habilidades := []Habilidad{habilidad1, habilidad2, habilidad3}
	datos := Datos{Nombre: "Chris", Edad: 20, Perfil: 1, Habilidades: habilidades}
	template := template.Must(template.ParseFiles("views/estructuras.html", utils.Frontend))
	template.Execute(response, datos)
}
