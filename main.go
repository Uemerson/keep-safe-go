package main

import (
	"log"
	"net/http"

	"github.com/Uemerson/keep-safe-go/keep"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	keep.Handler(r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
