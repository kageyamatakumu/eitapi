package db

import (
	"fmt"
	"log"
	"os"


	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDB() *gorm.DB {
	var url string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DATABASE"))

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Database connected")
	return db
}
