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

func IndexNoteService() []Note {
	return Notes
}

func ShowNoteService(uuid string) Note {
	var note Note
	for _, item := range Notes {
		if item.UUID == uuid {
			note = item
		}
	}
	return note
}

func DeleteNoteService(uuid string) {
	for index, note := range Notes {
		if note.UUID == uuid {
			Notes = append(Notes[:index], Notes[index+1:]...)
		}
	}
}
