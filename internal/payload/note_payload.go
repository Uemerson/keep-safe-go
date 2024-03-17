package payload

import (
	"time"
)

type NoteRequest struct {
	Text string `json:"text"`
}

type NoteResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Text      string    `json:"text"`
}

func NewNoteResponde(id string, createdAt, updatedAt time.Time, text string) *NoteResponse {
	return &NoteResponse{
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Text:      text,
	}
}
