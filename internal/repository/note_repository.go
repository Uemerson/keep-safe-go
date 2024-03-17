package repository

import (
	"context"
	"time"

	"github.com/Uemerson/keep-safe-go/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
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

func (nr *NoteRepository) GetNotes() ([]*entity.NoteEntity, error) {
	notesCollection := nr.db.Database("keepsafe").Collection("notes")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := notesCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var results []*entity.NoteEntity
	for cur.Next(ctx) {
		var result *entity.NoteEntity
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
