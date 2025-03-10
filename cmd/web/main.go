package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", getSnippetView)
	mux.HandleFunc("GET /snippet/create", getSnippetCreate)
	mux.HandleFunc("POST /snippet/create", postSnippetCreate)

	err := http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Server started at http://localhost:4000")

}
