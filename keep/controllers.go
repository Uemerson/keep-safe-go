package keep

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var Notes []Note

func Index(w http.ResponseWriter, r *http.Request) {
	notes := IndexNoteService()
	json.NewEncoder(w).Encode(notes)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	note := ShowNoteService(uuid)
	json.NewEncoder(w).Encode(note)
}

func Create(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var note Note
	json.Unmarshal(reqBody, &note)
	note = CreateNoteService(note)
	json.NewEncoder(w).Encode(note)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	DeleteNoteService(uuid)
}

func Update(w http.ResponseWriter, r *http.Request) {
}

func Patch(w http.ResponseWriter, r *http.Request) {
}
