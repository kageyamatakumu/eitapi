package repository

import (
	"backend/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_wordBookRepository_GetAllWordBook(t *testing.T) {
	// テスト用DBを作成する
	db := SetupTestDB()
	// テスト後にDBをクリーンアップする
	defer func() {
		_ = db.Migrator().DropTable(
			&model.Admin{},
			&model.Player{},
			&model.WordBook{},
			&model.Word{},
			&model.Genre{},
			&model.Difficulty{})
	}()

	repo := NewWordBookRepository(db)

	// テストデータ準備
	admin := model.Admin{Email: "test@example.com", Password: "password123"}
	db.Create(&admin)
	genre := model.Genre{GenreName: "test genre"}
	db.Create(&genre)
	difficulty := model.Difficulty{DifficultyLevel: 1}
	db.Create(&difficulty)
	wordBooks := []model.WordBook{
		{Title: "test word book 1", Description: "test description 1", AdminId: admin.ID, GenreId: genre.ID, DifficultyId: difficulty.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Title: "test word book 2", Description: "test description 2", AdminId: admin.ID, GenreId: genre.ID, DifficultyId: difficulty.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	db.Create(&wordBooks)

	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"success", false},
		{"no word books", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "no word books" {
				// テストデータを削除
				db.Where("1 = 1").Delete(&model.WordBook{})
			}
			var gotWordBooks []model.WordBook
			err := repo.GetAllWordBooks(&gotWordBooks)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllWordBook() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if tt.name == "success" {
					assert.Equal(t, len(wordBooks), len(gotWordBooks))
					for i, wordBook := range wordBooks {
						assert.Equal(t, wordBook.Title, gotWordBooks[i].Title)
						assert.Equal(t, wordBook.Description, gotWordBooks[i].Description)
						assert.Equal(t, wordBook.AdminId, gotWordBooks[i].AdminId)
						assert.Equal(t, wordBook.GenreId, gotWordBooks[i].GenreId)
						assert.Equal(t, wordBook.DifficultyId, gotWordBooks[i].DifficultyId)
					}
				} else if tt.name == "no word books" {
					assert.Equal(t, 0, len(gotWordBooks))
				}
			}
		})
	}
}

func Test_wordBookRepository_CreateWordBook(t *testing.T) {
	// テスト用DBを作成する
	db := SetupTestDB()
	// テスト後にDBをクリーンアップする
	defer func() {
		_ = db.Migrator().DropTable(
			&model.Admin{},
			&model.Player{},
			&model.WordBook{},
			&model.Word{},
			&model.Genre{},
			&model.Difficulty{})
	}()

	repo := NewWordBookRepository(db)

	// テストデータ準備
	admin := model.Admin{Email: "test@example.com", Password: "password123"}
	db.Create(&admin)
	genre := model.Genre{GenreName: "test genre"}
	db.Create(&genre)
	difficulty := model.Difficulty{DifficultyLevel: 1}
	db.Create(&difficulty)

	tests := []struct {
		name     string
		wordBook *model.WordBook
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"success", &model.WordBook{Title: "success test word book", Description: "test description", AdminId: admin.ID, GenreId: genre.ID, DifficultyId: difficulty.ID}, false},
		{"no description", &model.WordBook{Title: "no description test word book", AdminId: admin.ID, GenreId: genre.ID, DifficultyId: difficulty.ID}, false},
		{"admin not found", &model.WordBook{Title: "admin not found test word book", Description: "test description", AdminId: 999, GenreId: genre.ID, DifficultyId: difficulty.ID}, true},
		{"genre not found", &model.WordBook{Title: "genre not found test word book", Description: "test description", AdminId: admin.ID, GenreId: 999, DifficultyId: difficulty.ID}, true},
		{"difficulty not found", &model.WordBook{Title: "difficulty not found test word book", Description: "test description", AdminId: admin.ID, GenreId: genre.ID, DifficultyId: 999}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.CreateWordBook(tt.wordBook)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWordBook() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				var gotWordBook model.WordBook
				result := db.Where("title = ?", tt.wordBook.Title).First(&gotWordBook)
				assert.NoError(t, result.Error)
				assert.Equal(t, tt.wordBook.Title, gotWordBook.Title)
				if tt.wordBook.Description == "" {
					assert.Equal(t, "未設定", gotWordBook.Description)
				} else {
					assert.Equal(t, tt.wordBook.Description, gotWordBook.Description)
				}
				assert.Equal(t, tt.wordBook.AdminId, gotWordBook.AdminId)
				assert.Equal(t, tt.wordBook.GenreId, gotWordBook.GenreId)
				assert.Equal(t, tt.wordBook.DifficultyId, gotWordBook.DifficultyId)
			}
		})
	}
}

func Test_wordBookRepository_DeleteWordBook(t *testing.T) {
	// テスト用DBを作成する
	db := SetupTestDB()
	// テスト後にDBをクリーンアップする
	defer func() {
		_ = db.Migrator().DropTable(
			&model.Admin{},
			&model.Player{},
			&model.WordBook{},
			&model.Word{},
			&model.Genre{},
			&model.Difficulty{})
	}()

	repo := NewWordBookRepository(db)

	// テストデータ準備
	admin := model.Admin{Email: "test@example.com", Password: "password123"}
	db.Create(&admin)
	genre := model.Genre{GenreName: "test genre"}
	db.Create(&genre)
	difficulty := model.Difficulty{DifficultyLevel: 1}
	db.Create(&difficulty)
	wordBooks := []model.WordBook{
		{Title: "test word book 1", Description: "test description 1", AdminId: admin.ID, GenreId: genre.ID, DifficultyId: difficulty.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Title: "test word book 2", Description: "test description 2", AdminId: admin.ID, GenreId: genre.ID, DifficultyId: difficulty.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	db.Create(&wordBooks)

	tests := []struct {
		name       string
		wordBookId uint
		wantErr    bool
	}{
		// TODO: Add test cases.
		{"success", wordBooks[0].ID, false},
		{"not found", 999, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.DeleteWordBook(tt.wordBookId)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteWordBook() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				// 削除後のデータ検証
				var gotWordBook model.WordBook
				result := db.Where("id = ?", tt.wordBookId).First(&gotWordBook)
				assert.Error(t, result.Error) // データが存在しないことを検証
			} else {
				if tt.name == "not found" {
					assert.Equal(t, gorm.ErrRecordNotFound, err)
					assert.Equal(t, gorm.ErrRecordNotFound.Error(), err.Error())
				}
			}
		})
	}
}
