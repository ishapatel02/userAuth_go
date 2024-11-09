package database

import (
	"context"
	"fmt"
	"go_user_authentication/configuration"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var MongoClient *mongo.Client
// var UserCollection *mongo.Collection

// func ConnectMongoDB() {
// 	fmt.Println(configuration.AppConfig.DB_URI)
// 	clientOptions := options.Client().ApplyURI(configuration.AppConfig.DB_URI)
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to MongoDB: %v", err)
// 	}

// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to ping MongoDB: %v", err)
// 	}

// 	// MongoClient = client
// 	UserCollection = client.Database("go_auth").Collection("users")
// 	log.Println("Connected to MongoDB")
// }

type MongoDB struct {
	Client *mongo.Client
}

var instance *MongoDB
var once sync.Once

func ConnectMongoDB() *MongoDB {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI(configuration.AppConfig.DB_URI)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connected to MongoDB!")

		instance = &MongoDB{
			Client: client,
		}

	})
	return instance
}

func GetCollection(databaseName, collectionName string) *mongo.Collection {
	return instance.Client.Database(databaseName).Collection(collectionName)
}
