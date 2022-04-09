package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Uemerson/keep-safe-go/keep"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	log.Println("connected to mongodb!")
	r := mux.NewRouter()
	keep.Handler(r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
