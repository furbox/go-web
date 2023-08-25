package rutas

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(response, "Hello World")
	template, errorTemplate := template.ParseFiles("views/home.html")
	if errorTemplate != nil {
		panic(errorTemplate)
	}
	template.Execute(response, nil)
}
func Nosotros(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(response, "Nosotros - Hola Mundo")
	template, errorTemplate := template.ParseFiles("views/nosotros.html")
	if errorTemplate != nil {
		panic(errorTemplate)
	}
	template.Execute(response, nil)
}
func Params(response http.ResponseWriter, request *http.Request) {
	/*
		vars := mux.Vars(request)
		fmt.Fprintf(response, "Parametros: "+vars["id"]+" "+vars["slug"])
	*/
	template, errorTemplate := template.ParseFiles("views/params.html")
	if errorTemplate != nil {
		panic(errorTemplate)
	}
	vars := mux.Vars(request)
	data := map[string]string{
		"id":   vars["id"],
		"slug": vars["slug"],
	}
	template.Execute(response, data)
}
func Queries(response http.ResponseWriter, request *http.Request) {
	//request.URL
	//request.URL.RawQuery
	/*
		query := request.URL.Query()
		fmt.Fprintf(response, "Queries: "+query.Get("id")+" "+query.Get("name"))
	*/
	template, errorTemplate := template.ParseFiles("views/queries.html")
	if errorTemplate != nil {
		panic(errorTemplate)
	}
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
	template, errorTemplate := template.ParseFiles("views/estructuras.html")
	if errorTemplate != nil {
		panic(errorTemplate)
	}
	template.Execute(response, datos)
}
