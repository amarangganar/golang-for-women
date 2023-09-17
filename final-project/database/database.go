package database

import (
	"final_project/model"
	"final_project/repository"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() (*repository.Database, error) {
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	dbuser := os.Getenv("DB_USERNAME")
	dbpassword := os.Getenv("DB_PASSWORD")

	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbport, dbname)
	db, err := gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err := db.Debug().AutoMigrate(&model.Admin{}, &model.Product{}, &model.Variant{}); err != nil {
		return nil, err
	}

	return &repository.Database{DB: db}, nil
}
