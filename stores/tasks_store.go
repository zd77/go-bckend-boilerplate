package stores

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskStore interface {
	GetTasks() ([]string, error)
}

type MongoTaskStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewTaskStore(client *mongo.Client) *MongoTaskStore {
	return &MongoTaskStore{
		client:     client,
		collection: client.Database("").Collection(""),
	}
}

func (ts *MongoTaskStore) GetTasks() ([]string, error) {
	robins := []string{"Dick Grayson", "Jason Todd", "Tim Drake", "Damian Wayne"}
	return robins, nil
}
