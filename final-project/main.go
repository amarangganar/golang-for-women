package main

import (
	"final_project/database"
	"final_project/router"
	"fmt"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cld, err = cloudinary.New()
	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
	}

	db, err := database.Init()
	if err != nil {
		panic(err)
	}

	router.Execute(db, cld).Run(os.Getenv("SERVER_PORT"))

	fmt.Println("Hello world!")
}
