package keep

import (
	"github.com/gorilla/mux"
)

func Handler(r *mux.Router) {
	keepRouter := r.PathPrefix("/keep").Subrouter()
	keepRouter.HandleFunc("/", Create).Methods("POST")
	keepRouter.HandleFunc("/", Index).Methods("GET")
	keepRouter.HandleFunc("/{uuid}", Show).Methods("GET")
	keepRouter.HandleFunc("/{uuid}", Delete).Methods("DELETE")
	keepRouter.HandleFunc("/{uuid}", Update).Methods("PUT")
	keepRouter.HandleFunc("/{uuid}", Patch).Methods("PATCH")
}
