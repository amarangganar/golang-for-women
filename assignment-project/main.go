package main

import (
	"assignmentproject/database"
	"assignmentproject/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	env_err := godotenv.Load()
	if env_err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Start()
	routes.Start().Run(os.Getenv("SERVER_PORT"))
}
