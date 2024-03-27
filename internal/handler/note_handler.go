package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Uemerson/keep-safe-go/internal/entity"
	"github.com/Uemerson/keep-safe-go/internal/exception"
	"github.com/Uemerson/keep-safe-go/internal/service"
)

type NoteHandler struct {
	ns *service.NoteService
}

func NewNoteHandler(ns *service.NoteService) *NoteHandler {
	return &NoteHandler{ns: ns}
}

func (nh *NoteHandler) AddNote(w http.ResponseWriter, r *http.Request) {
	var noteRequest entity.NoteEntity
	if err := json.NewDecoder(r.Body).Decode(&noteRequest); err != nil {
		err := exception.NewInternalServerError(err.Error())
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}
	note, err := nh.ns.AddNote(entity.NewNoteEntity(noteRequest.Text))
	if err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := json.NewEncoder(w).Encode(note); err != nil {
		err := exception.NewInternalServerError(err.Error())
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
	}
}

func (nh *NoteHandler) LoadNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := nh.ns.LoadNotes()
	if err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := json.NewEncoder(w).Encode(notes); err != nil {
		err := exception.NewInternalServerError(err.Error())
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
	}
}

func (nh *NoteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := nh.ns.DeleteNote(id); err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (nh *NoteHandler) LoadNoteById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	note, err := nh.ns.LoadNoteById(id)
	if err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := json.NewEncoder(w).Encode(note); err != nil {
		err := exception.NewInternalServerError(err.Error())
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
	}
}

func (nh *NoteHandler) SaveNote(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var noteRequest entity.NoteEntity
	if err := json.NewDecoder(r.Body).Decode(&noteRequest); err != nil {
		err := exception.NewInternalServerError(err.Error())
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}
	note, err := nh.ns.SaveNote(id, &noteRequest)
	if err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := json.NewEncoder(w).Encode(note); err != nil {
		err := exception.NewInternalServerError(err.Error())
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
	}
}
