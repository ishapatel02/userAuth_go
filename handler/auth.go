package handler

import (
	"encoding/json"
	// "go_user_authentication/handler"
	"go_user_authentication/models"
	"go_user_authentication/services"
	"go_user_authentication/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// login Logic.
func Authenticate(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&creds)

	user, err := services.FindUserByUsername(creds.Username)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Username, user.Role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func Register(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	json.NewDecoder(r.Body).Decode(&creds)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	user := models.User{
		// ID:           primitive.NewObjectID(),
		Username:     creds.Username,
		PasswordHash: string(passwordHash),
		Role:         creds.Role,
	}

	err = services.CreateUser(&user)
	to := []string{"ishapatel2021@gmail.com"}
	subject := "Hello from Go!"
	body := "This is a test email sent from a Go program."

	SendEmail(subject, body, to)

	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created"})
}
