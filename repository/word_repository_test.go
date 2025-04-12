package repository

import (
	"backend/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_wordRepository_GetAllWordsForWordBook(t *testing.T) {
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
			&model.Difficulty{},
		)
	}()

	repo := NewWordRepository(db)

	// テストデータ準備
	admin := model.Admin{Email: "test@example.com", Password: "password123"}
	db.Create(&admin)
	genre := model.Genre{GenreName: "test genre"}
	db.Create(&genre)
	difficulty := model.Difficulty{DifficultyLevel: model.Easy}
	db.Create(&difficulty)
	wordBooks := []model.WordBook{
		{
			Title:        "test word book 1",
			Description:  "test description 1",
			AdminId:      admin.ID,
			GenreId:      genre.ID,
			DifficultyId: difficulty.ID,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Title:        "test word book 2",
			Description:  "test description 2",
			AdminId:      admin.ID,
			GenreId:      genre.ID,
			DifficultyId: difficulty.ID,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}
	db.Create(&wordBooks)
	words := []model.Word{
		{
			EnglishWord:         "test",
			JapaneseTranslation: "テスト",
			Pronunciation:       "test",
			ExampleSentence:     "test",
			WordBookId:          wordBooks[0].ID,
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
		},
		{
			EnglishWord:         "test2",
			JapaneseTranslation: "テスト2",
			Pronunciation:       "test2",
			ExampleSentence:     "test2",
			WordBookId:          wordBooks[0].ID,
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
		},
	}
	db.Create(&words)

	tests := []struct {
		name       string
		wordBookId uint
		wantWords  *[]model.Word
		wantErr    bool
	}{
		{"success", wordBooks[0].ID, &words, false},
		{"no words", wordBooks[1].ID, &words, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotWords []model.Word
			err := repo.GetAllWordsForWordBook(tt.wordBookId, &gotWords)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllWordsForWordBook() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if tt.name == "success" {
					assert.Equal(t, len(*tt.wantWords), len(gotWords))
					for i, word := range *tt.wantWords {
						assert.Equal(t, word.EnglishWord, gotWords[i].EnglishWord)
						assert.Equal(t, word.JapaneseTranslation, gotWords[i].JapaneseTranslation)
						assert.Equal(t, word.Pronunciation, gotWords[i].Pronunciation)
						assert.Equal(t, word.ExampleSentence, gotWords[i].ExampleSentence)
						assert.Equal(t, word.WordBookId, gotWords[i].WordBookId)
					}
				} else if tt.name == "no words" {
					assert.Equal(t, 0, len(gotWords))
					var count int64
					db.Model(&model.Word{}).Where("word_book_id = ?", tt.wordBookId).Count(&count)
					assert.Equal(t, int64(0), count)
				}
			}
		})
	}
}

func Test_wordRepository_GetWordById(t *testing.T) {
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
			&model.Difficulty{},
		)
	}()

	repo := NewWordRepository(db)

	// テストデータ準備
	admin := model.Admin{Email: "test@example.com", Password: "password123"}
	db.Create(&admin)
	genre := model.Genre{GenreName: "test genre"}
	db.Create(&genre)
	difficulty := model.Difficulty{DifficultyLevel: model.Easy}
	db.Create(&difficulty)
	wordBook := model.WordBook{Title: "test word book 1", Description: "test description 1", AdminId: admin.ID, GenreId: genre.ID, DifficultyId: difficulty.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	db.Create(&wordBook)
	word := model.Word{EnglishWord: "test", JapaneseTranslation: "テスト", Pronunciation: "test", ExampleSentence: "test", WordBookId: wordBook.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	db.Create(&word)

	tests := []struct {
		name     string
		wordId   uint
		wantWord model.Word
		wantErr  bool
	}{
		{"success", word.ID, word, false},
		{"no words", 999, model.Word{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotWord model.Word
			err := repo.GetWordById(tt.wordId, &gotWord)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWordById() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if tt.name == "success" {
					assert.Equal(t, tt.wantWord.EnglishWord, gotWord.EnglishWord)
					assert.Equal(t, tt.wantWord.JapaneseTranslation, gotWord.JapaneseTranslation)
					assert.Equal(t, tt.wantWord.Pronunciation, gotWord.Pronunciation)
					assert.Equal(t, tt.wantWord.ExampleSentence, gotWord.ExampleSentence)
					assert.Equal(t, tt.wantWord.WordBookId, gotWord.WordBookId)
				} else if tt.name == "no words" {
					var count int64
					db.Model(&model.Word{}).Where("id = ?", tt.wordId).Count(&count)
					assert.Equal(t, int64(0), count)
				}
			}
		})
	}
}

func Test_wordRepository_CreateMultipleWords(t *testing.T) {
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
			&model.Difficulty{},
		)
	}()

	repo := NewWordRepository(db)

	// テストデータ準備
	admin := model.Admin{Email: "test@example.com", Password: "password123"}
	db.Create(&admin)
	genre := model.Genre{GenreName: "test genre"}
	db.Create(&genre)
	difficulty := model.Difficulty{DifficultyLevel: model.Easy}
	db.Create(&difficulty)
	wordBooks := []model.WordBook{
		{
			Title:        "test word book 1",
			Description:  "test description 1",
			AdminId:      admin.ID,
			GenreId:      genre.ID,
			DifficultyId: difficulty.ID,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Title:        "test word book 2",
			Description:  "test description 2",
			AdminId:      admin.ID,
			GenreId:      genre.ID,
			DifficultyId: difficulty.ID,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}
	db.Create(&wordBooks)

	tests := []struct {
		name    string
		words   *[]model.Word
		wantErr bool
	}{
		{
			"success",
			&[]model.Word{
				{
					EnglishWord:         "test",
					JapaneseTranslation: "テスト",
					Pronunciation:       "test",
					ExampleSentence:     "test",
					WordBookId:          wordBooks[0].ID,
					CreatedAt:           time.Now(),
					UpdatedAt:           time.Now(),
				},
			},
			false,
		},
		{
			"success multiple words",
			&[]model.Word{
				{
					EnglishWord:         "test2",
					JapaneseTranslation: "テスト2",
					Pronunciation:       "test2",
					ExampleSentence:     "test2",
					WordBookId:          wordBooks[1].ID,
					CreatedAt:           time.Now(),
					UpdatedAt:           time.Now(),
				},
				{
					EnglishWord:         "test3",
					JapaneseTranslation: "テスト3",
					Pronunciation:       "test3",
					ExampleSentence:     "test3",
					WordBookId:          wordBooks[1].ID,
					CreatedAt:           time.Now(),
					UpdatedAt:           time.Now(),
				},
			},
			false,
		},
		{
			"word book not found",
			&[]model.Word{
				{
					EnglishWord:         "test4",
					JapaneseTranslation: "テスト4",
					Pronunciation:       "test4",
					ExampleSentence:     "test4",
					WordBookId:          999,
					CreatedAt:           time.Now(),
					UpdatedAt:           time.Now(),
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.CreateMultipleWords(tt.words)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMultipleWords() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				for _, word := range *tt.words {
					var gotWord model.Word
					result := db.Where("english_word = ?", word.EnglishWord).First(&gotWord)
					assert.NoError(t, result.Error)
					assert.Equal(t, word.EnglishWord, gotWord.EnglishWord)
					assert.Equal(t, word.JapaneseTranslation, gotWord.JapaneseTranslation)
					assert.Equal(t, word.Pronunciation, gotWord.Pronunciation)
					assert.Equal(t, word.ExampleSentence, gotWord.ExampleSentence)
					assert.Equal(t, word.WordBookId, gotWord.WordBookId)
				}
			} else {
				if tt.name == "word book not found" {
					assert.Error(t, err)
				}
			}
		})
	}
}

func Test_wordRepository_UpdateWord(t *testing.T) {
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
			&model.Difficulty{},
		)
	}()

	repo := NewWordRepository(db)

	// テストデータ準備
	admin := model.Admin{Email: "test@example.com", Password: "password123"}
	db.Create(&admin)
	genre := model.Genre{GenreName: "test genre"}
	db.Create(&genre)
	difficulty := model.Difficulty{DifficultyLevel: model.Easy}
	db.Create(&difficulty)
	wordBook := model.WordBook{Title: "test word book 1", Description: "test description 1", AdminId: admin.ID, GenreId: genre.ID, DifficultyId: difficulty.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	db.Create(&wordBook)
	word := model.Word{EnglishWord: "test", JapaneseTranslation: "テスト", Pronunciation: "test", ExampleSentence: "test", WordBookId: wordBook.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	db.Create(&word)

	tests := []struct {
		name    string
		word    *model.Word
		wordId  uint
		wantErr bool
	}{
		{"success", &model.Word{EnglishWord: "update test", JapaneseTranslation: "アップデート テスト", Pronunciation: "update test", ExampleSentence: "update test"}, word.ID, false},
		{"not found", &model.Word{EnglishWord: "not found test", JapaneseTranslation: "ノット ファウンド テスト", Pronunciation: "not found test", ExampleSentence: "not found update test"}, 999, true},
		{"partial update", &model.Word{EnglishWord: "partial update test", JapaneseTranslation: "アップデート テスト", Pronunciation: "update test", ExampleSentence: "update test"}, word.ID, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.UpdateWord(tt.word, tt.wordId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateWord() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				var updatedWord model.Word
				result := db.Where("id = ?", tt.wordId).First(&updatedWord)
				assert.NoError(t, result.Error)
				assert.Equal(t, tt.word.EnglishWord, updatedWord.EnglishWord)
				assert.Equal(t, tt.word.JapaneseTranslation, updatedWord.JapaneseTranslation)
				assert.Equal(t, tt.word.Pronunciation, updatedWord.Pronunciation)
				assert.Equal(t, tt.word.ExampleSentence, updatedWord.ExampleSentence)
			} else {
				if err == gorm.ErrRecordNotFound {
					assert.Equal(t, gorm.ErrRecordNotFound, err)
					assert.Equal(t, gorm.ErrRecordNotFound.Error(), err.Error())
				}
			}
		})
	}
}

func Test_wordRepository_DeleteWord(t *testing.T) {
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
			&model.Difficulty{},
		)
	}()

	repo := NewWordRepository(db)

	// テストデータ準備
	admin := model.Admin{Email: "test@example.com", Password: "password123"}
	db.Create(&admin)
	genre := model.Genre{GenreName: "test genre"}
	db.Create(&genre)
	difficulty := model.Difficulty{DifficultyLevel: model.Easy}
	db.Create(&difficulty)
	wordBook := model.WordBook{Title: "test word book 1", Description: "test description 1", AdminId: admin.ID, GenreId: genre.ID, DifficultyId: difficulty.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	db.Create(&wordBook)
	word := model.Word{EnglishWord: "test", JapaneseTranslation: "テスト", Pronunciation: "test", ExampleSentence: "test", WordBookId: wordBook.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	db.Create(&word)

	tests := []struct {
		name    string
		wordId  uint
		wantErr bool
	}{
		// TODO: Add test cases.
		{"success", word.ID, false},
		{"not found", 999, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.DeleteWord(tt.wordId)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteWord() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				var deletedWord model.Word
				result := db.Where("id = ?", tt.wordId).First(&deletedWord)
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
