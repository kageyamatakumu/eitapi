package model

import "time"

type Admin struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Email     string    `json:"email" gorm:"not null;unique type:VARCHAR(255)"`
	Password  string    `json:"password" gorm:"not null type:VARCHAR(255)"`
	UserName  string    `json:"user_name" gorm:"default: '未設定' type:VARCHAR(255)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
