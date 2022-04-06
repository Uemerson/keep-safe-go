package keep

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var Notes []Note

func Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Notes)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["uuid"]
	for _, note := range Notes {
		if note.UUID == key {
			json.NewEncoder(w).Encode(note)
		}
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
}

func Delete(w http.ResponseWriter, r *http.Request) {
}

func Update(w http.ResponseWriter, r *http.Request) {
}

func Patch(w http.ResponseWriter, r *http.Request) {
}
