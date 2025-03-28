package model

import "time"

type Word struct {
	ID                  uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	EnglishWord         string    `json:"english_word" gorm:"not null;type:VARCHAR(255)"`
	JapaneseTranslation string    `json:"japanese_translation" gorm:"not null;type:VARCHAR(255)"`
	Pronunciation       string    `json:"pronunciation" gorm:"not null;type:VARCHAR(255)"`
	ExampleSentence     string    `json:"example_sentence" gorm:"default:'未設定';type:TEXT"`
	WordBookId          uint      `json:"word_book_id" gorm:"not null"`
	CreatedAt           time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt           time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Association
	WordBook WordBook `json:"word_book"`
}

type WordRes struct {
	ID                  uint   `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	EnglishWord         string `json:"english_word" gorm:"not null;type:VARCHAR(255)"`
	JapaneseTranslation string `json:"japanese_translation" gorm:"not null;type:VARCHAR(255)"`
	Pronunciation       string `json:"pronunciation" gorm:"not null;type:VARCHAR(255)"`
	ExampleSentence     string `json:"example_sentence" gorm:"default:'未設定';type:TEXT"`
	WordBookId          uint   `json:"word_book_id" gorm:"not null"`
}
