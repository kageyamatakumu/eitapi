package main

import (
	"backend/db"
	"backend/model"
	"log"
	"fmt"

	"gorm.io/gorm"
)

func main() {
	dbConn := db.CreateDB()

	defer fmt.Println("Successfully Migrated!")
	defer db.CloseDB(dbConn)

	err := dbConn.AutoMigrate(
		&model.Admin{},
		&model.Player{},
		&model.WordBook{},
		&model.Word{},
		&model.Genre{},
		&model.Difficulty{},
	)
	if err != nil {
		log.Fatalf("マイグレーションエラー: %v", err)
	}

	seed(dbConn)
}

func seed(db *gorm.DB) {
	admin := model.Admin{
		Email:    "email@example.com",
		Password: "$2a$10$xDlV8e5VAvLjNJNZ6gWHrOFoha.Xkdhk.nPq.hIOXLJPM9Ey3AuYm",
		UserName: "eitapi",
	}
	difficulty := model.Difficulty{
		DifficultyLevel: 0,
	}
	genre := model.Genre{
		GenreName: "初回マイグレーション",
	}
	wordBook := model.WordBook{
		AdminId:      1,
		GenreId:      1,
		DifficultyId: 1,
		Title:        "英単語帳（初回マイグレーション）",
		Description:  "マイグレーション実施時に自動で作成される英単語帳",
	}
	word := model.Word{
		WordBookId:          1,
		EnglishWord:         "apple",
		JapaneseTranslation: "りんご",
		Pronunciation:       "æpl",
		ExampleSentence:     "This is an apple.",
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Printf("admin seed error: %v", err)
	}
	if err := db.Create(&difficulty).Error; err != nil {
		log.Printf("difficulty seed error: %v", err)
	}
	if err := db.Create(&genre).Error; err != nil {
		log.Printf("genre seed error: %v", err)
	}
	if err := db.Create(&wordBook).Error; err != nil {
		log.Printf("word_book_id seed error: %v", err)
	}
	if err := db.Create(&word).Error; err != nil {
		log.Printf("word seed error: %v", err)
	}
}
