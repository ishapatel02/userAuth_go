package main

import (
	"context"
	"go_user_authentication/configuration"
	"go_user_authentication/database"
	"go_user_authentication/handler"
	"go_user_authentication/middleware"
	"go_user_authentication/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// var UserCollection *mongo.Collection

func main() {

	configuration.LoadConfig()

	// Connect to MongoDB and get the MongoDB instance
	mongoDB := database.ConnectMongoDB()
	handler.InitializeStripe(configuration.AppConfig.StripeKey)

	// Use the client from the MongoDB struct
	client := mongoDB.Client

	// Defer the disconnection of the MongoDB client
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// Initialize routes
	router := mux.NewRouter()

	router.Handle("/admin", middleware.Authorize("ADMIN")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome Admin!"))
	})))

	// Define the /user route with authorization middleware
	router.Handle("/user", middleware.Authorize("ADMIN", "USER")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome User!"))
	})))

	routes.AuthenticationRoutes(router)

	// Start the server
	http.ListenAndServe(":"+configuration.AppConfig.ServerPort, router)
}
