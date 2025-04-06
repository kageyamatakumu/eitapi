package model

import "time"

type DifficultyLevel uint

const (
	Easy DifficultyLevel = iota + 1
	Medium
	Hard
)

type Difficulty struct {
	ID              uint            `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	DifficultyLevel DifficultyLevel `json:"difficulty_level" gorm:"not null"`
	CreatedAt       time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
}

type DifficultyResponse struct {
	ID              uint            `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	DifficultyLevel DifficultyLevel `json:"difficulty_level" gorm:"not null"`
}
