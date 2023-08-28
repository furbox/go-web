package main

import (
	"fmt"
	"go_web/rutas"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	mux := mux.NewRouter()
	//rutas
	mux.HandleFunc("/", rutas.Home)
	mux.HandleFunc("/nosotros", rutas.Nosotros)
	mux.HandleFunc("/params/{id:.*}/{slug:.*}", rutas.Params)
	mux.HandleFunc("/queries", rutas.Queries)
	mux.HandleFunc("/estructuras", rutas.Estructuras)
	mux.HandleFunc("/forms", rutas.Forms)
	mux.HandleFunc("/forms-post", rutas.FormPost).Methods("POST")
	mux.HandleFunc("/uploads", rutas.Upload)
	mux.HandleFunc("/upload-post", rutas.UploadPost).Methods("POST")

	//Archivos estaticos hacia mux
	s := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	mux.PathPrefix("/public/").Handler(s)

	//error 404
	mux.NotFoundHandler = mux.NewRoute().HandlerFunc(rutas.Page404).GetHandler()

	//ejecucion del servidor
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}

	server := &http.Server{
		Addr:         "localhost:" + os.Getenv("PORT"),
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server running on port " + os.Getenv("PORT") + ", http://localhost:" + os.Getenv("PORT"))
	log.Fatal(server.ListenAndServe())
}
