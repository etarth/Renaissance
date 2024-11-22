package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Context  context.Context
	Client   *mongo.Client
	Database *mongo.Database
}

func ConnectToMongoDB(uri string, dbName string) *MongoDB {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	log.Println("Connected to MongoDB")
	return &MongoDB{
		Context:  context.Background(),
		Client:   client,
		Database: client.Database(dbName),
	}
}

func (m *MongoDB) Disconnect() {
	err := m.Client.Disconnect(m.Context)
	if err != nil {
		log.Fatalf("Error while disconnecting MongoDB: %v", err)
	} else {
		log.Println("Disconnected from MongoDB")
	}
}
