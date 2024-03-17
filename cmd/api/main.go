package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Uemerson/keep-safe-go/internal/route"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var ctx = context.TODO()
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf(
			"mongodb://%s:%s@mongo:%s/",
			os.Getenv("MONGO_USERNAME"),
			os.Getenv("MONGO_PASSWORD"),
			os.Getenv("MONGO_PORT")))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	route.NoteRoute(mux, client)
	fmt.Println("Server is running on port " + os.Getenv("API_PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("API_PORT")), mux)
}
