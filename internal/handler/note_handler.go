package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Uemerson/keep-safe-go/internal/entity"
	"github.com/Uemerson/keep-safe-go/internal/exception"
	"github.com/Uemerson/keep-safe-go/internal/payload"
	"github.com/Uemerson/keep-safe-go/internal/service"
)

type NoteHandler struct {
	ns *service.NoteService
}

func NewNoteHandler(ns *service.NoteService) *NoteHandler {
	return &NoteHandler{ns: ns}
}

func (nh *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var noteRequest payload.NoteRequest
	if err := json.NewDecoder(r.Body).Decode(&noteRequest); err != nil {
		err := exception.NewInternalServerError(err.Error())
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}
	note, err := nh.ns.CreateNote(entity.NewNoteEntity(noteRequest.Text))
	if err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := json.NewEncoder(w).Encode(payload.NewNoteResponde(
		note.ID.Hex(), note.CreatedAt,
		note.UpdatedAt, note.Text)); err != nil {
		err := exception.NewInternalServerError(err.Error())
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
	}
}
