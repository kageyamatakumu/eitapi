package repository

import (
	"backend/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_difficultyRepository_GetAllDifficulties(t *testing.T) {
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

	repo := NewDifficultyRepository(db)

	// テストデータ準備
	difficulties := []model.Difficulty{
		{DifficultyLevel: 1},
		{DifficultyLevel: 2},
	}
	db.Create(&difficulties)

	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"success", false},
		{"no difficulty", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "no difficulty" {
				// テストデータを削除
				db.Where("1 = 1").Delete(&model.Difficulty{})
			}
			var gotDifficulties []model.Difficulty
			err := repo.GetAllDifficulties(&gotDifficulties)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllDifficulties() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if tt.name == "success" {
					assert.Equal(t, len(difficulties), len(gotDifficulties))
					for i, difficulty := range difficulties {
						assert.Equal(t, difficulty.DifficultyLevel, gotDifficulties[i].DifficultyLevel)
					}
				} else if tt.name == "no difficulties" {
					assert.Equal(t, 0, len(gotDifficulties))
				}
			}
		})
	}
}

func Test_difficultyRepository_CreateDifficulty(t *testing.T) {
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

	repo := NewDifficultyRepository(db)

	// 重複テスト用のデータ作成
	difficulty := model.Difficulty{DifficultyLevel: 2}
	db.Create(&difficulty)

	tests := []struct {
		name       string
		difficulty *model.Difficulty
		wantErr    bool
	}{
		// TODO: Add test cases.
		{"success", &model.Difficulty{DifficultyLevel: 1}, false},
		{"duplicate", &model.Difficulty{DifficultyLevel: 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.CreateDifficulty(tt.difficulty)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDifficulty() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				var createdDifficulty model.Difficulty
				result := db.Where("difficulty_level", tt.difficulty.DifficultyLevel).First(&createdDifficulty)
				assert.NoError(t, result.Error)
				assert.Equal(t, tt.difficulty.DifficultyLevel, createdDifficulty.DifficultyLevel)
			}
		})
	}
}

func Test_difficultyRepository_UpdateDifficulty(t *testing.T) {
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

	repo := NewDifficultyRepository(db)

	// テスト用のデータ作成
	difficulties := []model.Difficulty{{DifficultyLevel: 1}, {DifficultyLevel: 2}}
	for _, g := range difficulties {
		db.Create(&g)
	}

	tests := []struct {
		name         string
		difficulty   *model.Difficulty
		difficultyId uint
		wantErr      bool
	}{
		// TODO: Add test cases.
		{"success", &model.Difficulty{DifficultyLevel: 3}, 1, false},
		{"duplicate", &model.Difficulty{DifficultyLevel: 3}, 2, true},
		{"not found", &model.Difficulty{DifficultyLevel: 1}, 999, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.UpdateDifficulty(tt.difficulty, tt.difficultyId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateDifficulty() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				var updateDifficulty model.Difficulty
				result := db.Where("id = ?", tt.difficultyId).First(&updateDifficulty)
				assert.NoError(t, result.Error)
				assert.Equal(t, tt.difficulty.DifficultyLevel, updateDifficulty.DifficultyLevel)
			} else {
				if err == gorm.ErrRecordNotFound {
					assert.Equal(t, "record not found", err.Error())
				}
			}
		})
	}
}

func Test_difficultyRepository_DeleteDifficulty(t *testing.T) {
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

	repo := NewDifficultyRepository(db)

	// テスト用のデータ作成
	difficulties := []model.Difficulty{{DifficultyLevel: 1}, {DifficultyLevel: 2}}
	db.Create(difficulties)

	tests := []struct {
		name         string
		difficultyId uint
		wantErr      bool
	}{
		// TODO: Add test cases.
		{"success", difficulties[0].ID, false},
		{"difficulty not found", 999, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.DeleteDifficulty(tt.difficultyId)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteDifficulty() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				var gotDeleteDifficulty model.Difficulty
				result := db.Where("id = ?", tt.difficultyId).First(&gotDeleteDifficulty)
				assert.Error(t, result.Error)
				assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
			} else {
				// エラーメッセージの検証
				if tt.name == "difficulty not found" {
					assert.Equal(t, gorm.ErrRecordNotFound, err)
				}
			}
		})
	}
}
