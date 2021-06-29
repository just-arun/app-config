package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDbConnection(connectinString string) (client *mongo.Client, err error) {

	// Set client options
	clientOptions := options.Client().ApplyURI(connectinString)

	// Connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return
	}
	return
}
