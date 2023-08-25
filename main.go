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
