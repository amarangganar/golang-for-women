package main

import (
	"final_project/database"
	"final_project/router"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.Init()
	if err != nil {
		panic(err)
	}

	router.Execute(db).Run(os.Getenv("SERVER_PORT"))

	fmt.Println("Hello world!")
}
