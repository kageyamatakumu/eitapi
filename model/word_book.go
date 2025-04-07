package model

import "time"

type WordBook struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Title        string    `json:"title" gorm:"not null;type:VARCHAR(255)"`
	Description  string    `json:"description" gorm:"default:'未設定';type:TEXT"`
	AdminId      uint      `json:"admin_id" gorm:"not null"`
	GenreId      uint      `json:"genre_id" gorm:"not null"`
	DifficultyId uint      `json:"difficulty_id" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Association
	Admin      Admin      `json:"admin"`
	Genre      Genre      `json:"genre"`
	Difficulty Difficulty `json:"difficulty"`
}

type WordBookResponse struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Title        string    `json:"title" gorm:"not null;type:VARCHAR(255)"`
	Description  string    `json:"description" gorm:"default:'未設定';type:TEXT"`
	AdminId      uint      `json:"admin_id" gorm:"not null"`
	GenreId      uint      `json:"genre_id" gorm:"not null"`
	DifficultyId uint      `json:"difficulty_id" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type WordBookListResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Genre       Genre      `json:"genre"`
	Difficulty  Difficulty `json:"difficulty"`
}
