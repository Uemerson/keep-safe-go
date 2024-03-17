package route

import (
	"net/http"

	"github.com/Uemerson/keep-safe-go/internal/handler"
	"github.com/Uemerson/keep-safe-go/internal/repository"
	"github.com/Uemerson/keep-safe-go/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func NoteRoute(mux *http.ServeMux, client *mongo.Client) {
	noteRepository := repository.NewNoteRepository(client)
	noteService := service.NewNoteService(noteRepository)
	noteHandler := handler.NewNoteHandler(noteService)
	mux.HandleFunc("POST /", noteHandler.Create)
}
