package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client // Global variable

// DB is the MongoDB connection instance
var DB *mongo.Database

// ConnectDB establishes a connection to MongoDB
func ConnectDB() error {
	fmt.Println("ℹ️ Connecting to MongoDB...")

	// Get MongoDB URI from environment variable
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb+srv://preetheshvechoor4263:HuZRwxMk7nOWfmg7@cluster0.y1km5.mongodb.net/EmployeePortal"
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	// Ping MongoDB to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not connect to MongoDB: %v", err)
	}

	// Assign database to global `DB` variable
	DB = client.Database("EmployeePortal")

	fmt.Println("✅ Connected to MongoDB!")
	return nil
}

// DisconnectDB closes the MongoDB connection
func DisconnectDB() {
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := client.Disconnect(ctx); err != nil {
			log.Printf("❌ Error disconnecting from MongoDB: %v", err)
		} else {
			fmt.Println("✅ Disconnected from MongoDB!")
		}
	}
}

// GetCollection returns a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	if client == nil {
		fmt.Println("⏳ Waiting for MongoDB client to initialize...")
		time.Sleep(2 * time.Second) // Wait before retrying
	}
	if client == nil {
		log.Fatal("❌ MongoDB client is still nil! Exiting...")
	}
	fmt.Println("✅ Fetching collection:", collectionName)
	return DB.Collection(collectionName)
}
