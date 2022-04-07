package keep

import (
	"github.com/google/uuid"
)

func CreateNoteService(request Note) Note {
	note := request
	note.UUID = uuid.New().String()
	Notes = append(Notes, note)
	return note
}
