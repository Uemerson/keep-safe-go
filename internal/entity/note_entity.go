package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoteEntity struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Text      string             `bson:"text,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

func NewNoteEntity(text string) *NoteEntity {
	return &NoteEntity{
		ID:        primitive.NewObjectID(),
		Text:      text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
