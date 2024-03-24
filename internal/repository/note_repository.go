package repository

import (
	"context"
	"errors"

	"github.com/Uemerson/keep-safe-go/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NoteRepository struct {
	db *mongo.Client
}

func NewNoteRepository(db *mongo.Client) *NoteRepository {
	return &NoteRepository{db: db}
}

func (nr *NoteRepository) AddNote(note *entity.NoteEntity) (*mongo.InsertOneResult, error) {
	notesCollection := nr.db.Database("keepsafe").Collection("notes")
	result, err := notesCollection.InsertOne(context.TODO(), note)
	return result, err
}

func (nr *NoteRepository) LoadNotes() ([]*entity.NoteEntity, error) {
	notesCollection := nr.db.Database("keepsafe").Collection("notes")
	ctx := context.TODO()
	cur, err := notesCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
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

func (nr *NoteRepository) DeleteNote(id string) error {
	notesCollection := nr.db.Database("keepsafe").Collection("notes")
	idPrimitive, errPrimitive := primitive.ObjectIDFromHex(id)
	if errPrimitive != nil {
		return errPrimitive
	}
	_, err := notesCollection.DeleteOne(context.TODO(), bson.M{"_id": idPrimitive})
	if err != nil {
		return err
	}
	return nil
}

func (nr *NoteRepository) LoadNoteById(id string) (*entity.NoteEntity, error) {
	notesCollection := nr.db.Database("keepsafe").Collection("notes")
	idPrimitive, errPrimitive := primitive.ObjectIDFromHex(id)
	if errPrimitive != nil {
		return nil, errPrimitive
	}
	var res entity.NoteEntity
	err := notesCollection.FindOne(context.TODO(), bson.M{"_id": idPrimitive}).Decode(&res)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, nil
		default:
			return nil, err
		}
	}
	return &res, nil
}
