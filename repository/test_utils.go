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

	// マイグレーション実行
	if err := db.AutoMigrate(&model.WordBook{}, &model.Admin{}); err != nil { // 複数のモデルをマイグレートする場合
		panic("failed to migrate test database")
	}

	return db
}
