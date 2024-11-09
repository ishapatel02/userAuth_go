package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_URI     string
	ServerPort string
	StripeKey  string
}

var AppConfig Config

func LoadConfig() {
	// Load environment variables from .env file
	// err := godotenv.Load()
	err := godotenv.Load("C:/Users/ParNe/Documents/Isha/Go_UserAuth/.env")
	// if err != nil {
	//     log.Fatalf("Error loading .env file: %v", err)
	// }
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Assign values to the config struct
	AppConfig = Config{
		DB_URI:     os.Getenv("DB_URI"),
		ServerPort: os.Getenv("SERVER_PORT"),
		StripeKey:  os.Getenv("STRIPE_KEY"),
	}

	log.Println("Config loaded successfully")
}
