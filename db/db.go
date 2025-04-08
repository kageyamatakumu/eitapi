package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateDB() *gorm.DB {
	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DATABASE"))

	appEnv := os.Getenv("APP_ENV")
	var gormLogLevel logger.LogLevel

	if appEnv == "development" {
		gormLogLevel = logger.Info
	} else {
		gormLogLevel = logger.Silent
	}

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(gormLogLevel),
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Database connected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
