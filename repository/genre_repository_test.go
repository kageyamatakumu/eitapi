package repository

import (
	"backend/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_genreRepository_CreateGenre(t *testing.T) {
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

	repo := NewGenreRepository(db)

	tests := []struct {
		name    string
		genre   *model.Genre
		wantErr bool
	}{
		// TODO: Add test cases.
		{"success", &model.Genre{GenreName: "test"}, false},
		{"duplicate", &model.Genre{GenreName: "test"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.CreateGenre(tt.genre)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateGenre() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				var createdGenre model.Genre
				result := db.Where("genre_name = ?", tt.genre.GenreName).First(&createdGenre)
				assert.NoError(t, result.Error)
				assert.Equal(t, tt.genre.GenreName, createdGenre.GenreName)
			}
		})
	}
}

func Test_genreRepository_UpdateGenre(t *testing.T) {
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

	repo := NewGenreRepository(db)

	// テストデータ準備
	genres := []model.Genre{{GenreName: "test"}, {GenreName: "duplicate"}}
	for _, genre := range genres {
		db.Create(&genre)
	}

	tests := []struct {
		name    string
		genre   *model.Genre
		genreId uint
		wantErr bool
	}{
		// TODO: Add test cases.
		{"success", &model.Genre{GenreName: "updated"}, 1, false},
		{"not found", &model.Genre{GenreName: "not found"}, 999, true},
		{"duplicate", &model.Genre{GenreName: "duplicate"}, 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.UpdateGenre(tt.genre, tt.genreId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateGenre() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				var updatedGenre model.Genre
				result := db.Where("id = ?", tt.genreId).First(&updatedGenre)
				assert.NoError(t, result.Error)
				assert.Equal(t, tt.genre.GenreName, updatedGenre.GenreName)
			} else {
				if tt.name == "not found" {
					assert.Equal(t, gorm.ErrRecordNotFound, err)
					assert.Equal(t, gorm.ErrRecordNotFound.Error(), err.Error())
				}
			}
		})
	}
}

func Test_genreRepository_DeleteGenre(t *testing.T) {
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

	repo := NewGenreRepository(db)

	// テストデータ準備
	genre := model.Genre{GenreName: "test"}
	db.Create(&genre)

	tests := []struct {
		name    string
		genreId uint
		wantErr bool
	}{
		// TODO: Add test cases.
		{"success", genre.ID, false},
		{"not found", 999, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.DeleteGenre(tt.genreId)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteGenre() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				var deletedGenre model.Genre
				result := db.Where("id = ?", tt.genreId).First(&deletedGenre)
				assert.Error(t, result.Error)
				assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
			} else {
				if tt.name == "not found" {
					assert.Equal(t, gorm.ErrRecordNotFound, err)
					assert.Equal(t, gorm.ErrRecordNotFound.Error(), err.Error())
				}
			}
		})
	}
}
