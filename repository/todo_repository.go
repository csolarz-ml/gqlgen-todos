package repository

import (
	"context"
	"log"
	"os"

	"github.com/csolarz-ml/gqlgen-todos/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE   = "graphql"
	COLLECTION = "todos"
)

type TodoRepository interface {
	Save(todo *model.Todo) *model.Todo
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
	ctx := context.TODO()

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

func (db *database) Save(todo *model.Todo) *model.Todo {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	res, err := collection.InsertOne(context.TODO(), todo)

	if err != nil {
		log.Fatal(err)
	}

	todo.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return todo
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
