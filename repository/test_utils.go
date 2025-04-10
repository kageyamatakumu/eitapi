package repository

import (
	"backend/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// テスト用DBを作成
func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to open test database")
	}

	// 外部キー制約を有効化
	db.Exec("PRAGMA foreign_keys = ON")

	// マイグレーション実行
	if err := db.AutoMigrate(
		&model.Admin{},
		&model.Player{},
		&model.WordBook{},
		&model.Word{},
		&model.Genre{},
		&model.Difficulty{},
	); err != nil { // 複数のモデルをマイグレートする場合
		panic("failed to migrate test database")
	}

	return db
}
