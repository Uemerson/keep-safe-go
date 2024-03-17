package repository

import (
	"context"

	"github.com/Uemerson/keep-safe-go/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type NoteRepository struct {
	db *mongo.Client
}

func NewNoteRepository(db *mongo.Client) *NoteRepository {
	return &NoteRepository{db: db}
}

func (nr *NoteRepository) SaveNote(note *entity.NoteEntity) (*mongo.InsertOneResult, error) {
	notesCollection := nr.db.Database("keepsafe").Collection("notes")
	result, err := notesCollection.InsertOne(context.TODO(), note)
	return result, err
}
