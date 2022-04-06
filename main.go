package main

import (
	"log"
	"net/http"

	"github.com/Uemerson/keep-safe-go/keep"
	"github.com/gorilla/mux"
)

func main() {
	keep.Notes = []keep.Note{
		{UUID: "c131ab81-d939-4bb2-a5c2-aad5bbc06edd", Text: "One text"},
		{UUID: "c131ab81-d939-4bb2-a5c2-aad5bbc06edd", Text: "One text"},
	}

	r := mux.NewRouter()
	keep.Handler(r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
