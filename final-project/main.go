package main

import (
	"final_project/database"
	"final_project/router"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func main() {
	var cld, err = cloudinary.New()
	if err != nil {
		log.Fatalf("failed to intialize Cloudinary, %v", err)
	}

	db, err := database.Init()
	if err != nil {
		panic(err)
	}

	router.Execute(db, cld).Run(os.Getenv("SERVER_PORT"))
}
