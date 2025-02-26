package database

import (
	"context"
	"fiberpress-api/config"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	// Get MongoDB URI and DB name from config
	mongoURI := config.AppConfig.MongoURI
	dbName := config.AppConfig.MongoDBName

	// Create MongoDB connection options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("❌ MongoDB Connection Failed: ", err)
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ MongoDB Ping Failed: ", err)
	}

	fmt.Println("✅ Connected to MongoDB")

	// Select the database
	DB = client.Database(dbName)
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
