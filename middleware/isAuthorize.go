package middleware

import (
	"context"
	"fmt"
	"go_user_authentication/utils"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key") // Replace with your actual secret key

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type contextKey string

const roleContextKey = contextKey("ADMIN")

// Extracts the role from the JWT token in the Authorization header
func getRoleFromRequest(r *http.Request) string {
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		// If no Authorization header or it's not a Bearer token, return empty string (unauthorized)
		return ""
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		// If there's an error or the token is invalid, return empty string (unauthorized)
		log.Println("Invalid Token: ", err)
		return ""
	}

	fmt.Println("Role in claims:", claims.Role)

	return claims.Role
}
func Authorize(roles ...string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing token", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			_, err := utils.ParseJWT(tokenString)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			role := getRoleFromRequest(r)

			for _, allowedRole := range roles {
				fmt.Println(allowedRole)
				fmt.Println(role, "$$$$$$$$$$$$")
				if role == allowedRole {
					// If the role is allowed, store it in the request context
					ctx := context.WithValue(r.Context(), roleContextKey, role)
					// Pass the request to the next handler with the new context
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
			}

			// if !authorized {
			// 	http.Error(w, "Forbidden", http.StatusForbidden)
			// 	return
			// }
			http.Error(w, "Forbidden", http.StatusForbidden)

			// next(w, r)
		}
	}
}
