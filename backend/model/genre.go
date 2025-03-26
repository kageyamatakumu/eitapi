package model

import "time"

type Genre struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	GenreName string    `json:"genre_name" gorm:"not null;unique;type:VARCHAR(255)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
