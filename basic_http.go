package main

import (
	"log"
	"net/http"
)

func basic_http() {
	//mux := http.NewServeMux()

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("Hello World"))
	})
	print("Server running on port 8081, http://localhost:8081")
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}
