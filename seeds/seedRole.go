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

var roleCollection *mongo.Collection

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

	// Get the role collection
	roleCollection = client.Database("EmployeePortal").Collection("roles")
}

func seedRoleData() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define roles
	roles := []interface{}{
		models.Role{Name: "Employee"},
		models.Role{Name: "Manager"},
		models.Role{Name: "HR"},
		models.Role{Name: "Admin"},
		models.Role{Name: "Delivery Manager"},
		models.Role{Name: "Super Admin"},
		models.Role{Name: "Finance"},
	}

	// Delete existing roles
	_, err := roleCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Fatal("Error deleting existing roles:", err)
	}
	fmt.Println("Existing roles deleted.")

	// Insert new roles
	_, err = roleCollection.InsertMany(ctx, roles)
	if err != nil {
		log.Fatal("Error inserting roles:", err)
	}
	fmt.Println("Data seeded successfully!")
}

func main() {
	seedRoleData()
}
