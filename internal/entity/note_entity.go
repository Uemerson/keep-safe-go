package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoteEntity struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Text      string             `json:"text" bson:"text,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}

func NewNoteEntity(text string) *NoteEntity {
	return &NoteEntity{
		ID:        primitive.NewObjectID(),
		Text:      text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
