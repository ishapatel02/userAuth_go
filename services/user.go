package services

import (
	"context"
	"go_user_authentication/database"
	"go_user_authentication/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func FindUserByUsername(username string) (*models.User, error) {
	var user models.User

	userCollection := database.GetCollection("go_auth", "users")

	filter := bson.M{"username": username}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := userCollection.FindOne(ctx, filter).Decode(&user)
	return &user, err
}

func CreateUser(user *models.User) error {
	userCollection := database.GetCollection("go_auth", "users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := userCollection.InsertOne(ctx, user)
	// _, err := database.GetCollection()
	return err
}
