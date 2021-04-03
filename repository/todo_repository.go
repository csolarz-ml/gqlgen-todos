package repository

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/csolarz-ml/gqlgen-todos/graph/model"
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

func New() TodoRepository {

	MONGO_DB := os.Getenv("MONGO_DB")

	options := options.Client().ApplyURI(MONGO_DB).SetMaxPoolSize(100)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	client, err := mongo.Connect(ctx, options)

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
	cursor, err := collection.Find(context.TODO(), nil)

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
