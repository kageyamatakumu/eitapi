package model

import "time"

type WordBook struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Title        string    `json:"title" gorm:"not null;type:VARCHAR(255)"`
	Description  string    `json:"description" gorm:"default: '未設定';type:TEXT"`
	AdminId      uint      `json:"admin_id" gorm:"foreignKey:AdminId"`
	GenreId      uint      `json:"genre_id" gorm:"foreignKey:GenreId"`
	DifficultyId uint      `json:"difficulty_id" gorm:"foreignKey:DifficultyId"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
