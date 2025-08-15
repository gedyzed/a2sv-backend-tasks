package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func DbInit() *mongo.Database {

	
    // 1) Load the MongoDB URI from environment variable
    mongoURI := os.Getenv("MONGO_URI")
	fmt.Println(mongoURI)
    if mongoURI == "" {
        log.Fatal("MONGO_URI environment variable not set")
    }

    // 2) Configure the ServerAPI (required for Atlas)
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

    // 3) Connect to MongoDB Atlas
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, opts)
    if err != nil {
        log.Fatalf("MongoDB connection error: %v", err)
    }

    // 4) Ping to verify connection
    if err := client.Ping(ctx, nil); err != nil {
        log.Fatalf("MongoDB ping failed: %v", err)
    }

    log.Println("Successfully connected to MongoDB Atlas")

    // 5) Return the specific database
    return client.Database("production") // change to your DB name
}
