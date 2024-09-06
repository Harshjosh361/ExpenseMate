package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global variables for MongoDB client and collections
var (
	Client             *mongo.Client
	CollectionExpense  *mongo.Collection
	CollectionCategory *mongo.Collection
	CollectionUser     *mongo.Collection
)

// ConnectDb initializes the MongoDB connection and sets up collections
func ConnectDb() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := os.Getenv("MONGO_URI")
	if connectionString == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	// ensures err is reset
	err = nil
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Verify connection
	if err := Client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to MongoDB")

	// Initialize collections
	CollectionExpense = Client.Database("ExpenseMate").Collection("expenses")
	CollectionCategory = Client.Database("ExpenseMate").Collection("categories")
	CollectionUser = Client.Database("ExpenseMate").Collection("users")
}
