package stores

import "go.mongodb.org/mongo-driver/mongo"

type UserStore interface {
	GetUsers() ([]string, error)
}

type MongoUserStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client:     client,
		collection: client.Database("").Collection(""),
	}
}

func (us *MongoUserStore) GetUsers() ([]string, error) {
	theBoys := []string{"Butcher", "Mothers Milk", "Frenchman", "Female"}
	return theBoys, nil
}
