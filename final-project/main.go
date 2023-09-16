package main

import (
	"final_project/router"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	router.Execute()
	fmt.Println("Hello world!")
}
