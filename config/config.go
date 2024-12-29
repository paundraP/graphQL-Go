package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetDSN() string {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}
