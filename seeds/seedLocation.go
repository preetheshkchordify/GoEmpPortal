package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/preetheshkchordify/go-mongo-crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var locationCollection *mongo.Collection

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Get the location collection
	locationCollection = client.Database("EmployeePortal").Collection("locations")
}

func seedLocationData() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define Location
	locations := []interface{}{
		models.Location{Name: "TVM"},
		models.Location{Name: "Mohali"},
		models.Location{Name: "US"},
	}

	// Delete existing Locations
	_, err := locationCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Fatal("Error deleting existing locations:", err)
	}
	fmt.Println("Existing locations deleted.")

	// Insert new Locations
	_, err = locationCollection.InsertMany(ctx, locations)
	if err != nil {
		log.Fatal("Error inserting Locations:", err)
	}
	fmt.Println("Data seeded successfully!")
}

// func main() {
// 	seedLocationData()
// }
