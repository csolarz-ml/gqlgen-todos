package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/csolarz-ml/gqlgen-todos/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE   = "graphql"
	COLLECTION = "todos"
)

type TodoRepository interface {
	Save(todo *model.Todo)
	Find() []*model.Todo
}

type database struct {
	client *mongo.Client
}

func NewTodoRepository() TodoRepository {

	MONGO_DB := os.Getenv("MONGO_DB")

	if MONGO_DB == "" {
		MONGO_DB = "mongodb://localhost:27017/"
	}

	clientOptions := options.Client().ApplyURI(MONGO_DB)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
		log.Fatal(ctx.Err())
	}

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &database{
		client: client,
	}
}

func (db *database) Save(todo *model.Todo) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), todo)

	if err != nil {
		log.Fatal(err)
	}
}

func (db *database) Find() []*model.Todo {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.TODO())
	var result []*model.Todo
	for cursor.Next(context.TODO()) {
		var t *model.Todo
		err := cursor.Decode(&t)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, t)
	}

	return result
}
