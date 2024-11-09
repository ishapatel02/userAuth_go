package models

type User struct {
	Username     string `bson:"username"`
	PasswordHash string `bson:"password_hash"`
	Role         string `bson:"role"`
}
