package routes

import (
	"go_user_authentication/handler"

	"github.com/gorilla/mux" // or any router you are using
)

func AuthenticationRoutes(router *mux.Router) {

	// Define routes
	router.HandleFunc("/login", handler.Authenticate).Methods("POST")
	router.HandleFunc("/register", handler.Register).Methods("POST")

	// Add more routes here
}
