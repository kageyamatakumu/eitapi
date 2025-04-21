package usecase

import (
	"backend/model"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

// --- モックの定義 ---

type IWordBookRepositoryMock struct {
	mock.Mock
}

func (m *IWordBookRepositoryMock) GetAllWordBooks(wordBooks *[]model.WordBook) error {
	args := m.Called(wordBooks)
	return args.Error(0)
}

func (m *IWordBookRepositoryMock) CreateWordBook(wordBook *model.WordBook) error {
	args := m.Called(wordBook)
	return args.Error(0)
}

func (m *IWordBookRepositoryMock) DeleteWordBook(wordBookId uint) error {
	args := m.Called(wordBookId)
	return args.Error(0)
}

type IWordBookValidatorMock struct {
	mock.Mock
}

func (m *IWordBookValidatorMock) ValidateTitle(title string) error {
	args := m.Called(title)
	return args.Error(0)
}

// --- テスト ---

func Test_wordBookUsecase_GetAllWordBooks(t *testing.T) {

	tests := []struct {
		name      string
		mockSetup func(mockRepo *IWordBookRepositoryMock)
		want      []model.WordBookListResponse
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name: "正常系: 英単語帳リスト取得成功",
			mockSetup: func(mockRepo *IWordBookRepositoryMock) {
				mockRepo.On("GetAllWordBooks", mock.Anything).Run(func(args mock.Arguments) {
					ptr := args.Get(0).(*[]model.WordBook)
					*ptr = []model.WordBook{
						{ID: 1, Title: "簡単英単語", Description: "中学一年生の英単語", AdminId: 1, GenreId: 1, DifficultyId: 1},
						{ID: 2, Title: "少し難しい英単語", Description: "中学二年生の英単語", AdminId: 1, GenreId: 1, DifficultyId: 2},
					}
				}).Return(nil)
			},
			want: []model.WordBookListResponse{
				{ID: 1, Title: "簡単英単語", Description: "中学一年生の英単語", Genre: model.Genre{}, Difficulty: model.Difficulty{}},
				{ID: 2, Title: "少し難しい英単語", Description: "中学二年生の英単語", Genre: model.Genre{}, Difficulty: model.Difficulty{}},
			},
			wantErr: false,
		},
		{
			name: "異常系: リポジトリエラー発生",
			mockSetup: func(mockRepo *IWordBookRepositoryMock) {
				mockRepo.On("GetAllWordBooks", mock.Anything).Return(errors.New("db error"))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(IWordBookRepositoryMock)
			mockValidator := new(IWordBookValidatorMock)
			uc := NewWordBookUsecase(mockRepo, mockValidator)

			tt.mockSetup(mockRepo)

			got, err := uc.GetAllWordBooks()

			if (err != nil) != tt.wantErr {
				t.Errorf("unexpected error: got %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unexpected result: got %+v, want %+v", got, tt.want)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func Test_wordBookUsecase_CreateWordBook(t *testing.T) {

	createdAtStr := "2025-04-21 11:22:00.000000000 +0900 JST m=+0.000000001"
	updatedAtStr := "2025-04-21 11:22:00.000000000 +0900 JST m=+0.000000001"
	createdAt, _ := time.Parse(time.RFC3339Nano, createdAtStr)
	updatedAt, _ := time.Parse(time.RFC3339Nano, updatedAtStr)

	tests := []struct {
		name               string
		input              model.WordBook
		mockRepoSetup      func(mockRepo *IWordBookRepositoryMock)
		mockValidatorSetup func(mockValidator *IWordBookValidatorMock)
		want               model.WordBookResponse
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			name: "正常系: 英単語帳作成成功",
			input: model.WordBook{
				ID:           1,
				Title:        "簡単英単語",
				Description:  "中学一年生の英単語",
				AdminId:      1,
				GenreId:      1,
				DifficultyId: 1,
				CreatedAt:    createdAt,
				UpdatedAt:    updatedAt,
			},
			mockRepoSetup: func(mockRepo *IWordBookRepositoryMock) {
				mockRepo.On("CreateWordBook", mock.Anything).Return(nil)
			},
			mockValidatorSetup: func(mockValidator *IWordBookValidatorMock) {
				mockValidator.On("ValidateTitle", mock.Anything).Return(nil)
			},
			want: model.WordBookResponse{
				ID:           1,
				Title:        "簡単英単語",
				Description:  "中学一年生の英単語",
				AdminId:      1,
				GenreId:      1,
				DifficultyId: 1,
				CreatedAt:    createdAt,
				UpdatedAt:    updatedAt,
			},
			wantErr: false,
		},
		{
			name: "異常系: 英単語帳作成失敗",
			input: model.WordBook{
				ID:           1,
				Title:        "簡単英単語",
				Description:  "中学一年生の英単語",
				AdminId:      1,
				GenreId:      1,
				DifficultyId: 1,
				CreatedAt:    createdAt,
				UpdatedAt:    updatedAt,
			},
			mockRepoSetup: func(mockRepo *IWordBookRepositoryMock) {
				mockRepo.On("CreateWordBook", mock.Anything).Return(errors.New("db error"))
			},
			mockValidatorSetup: func(mockValidator *IWordBookValidatorMock) {
				mockValidator.On("ValidateTitle", mock.Anything).Return(nil)
			},
			want: model.WordBookResponse{},
			wantErr: true,
		},
		{
			name: "異常系: 不正なTitleでバリデーションエラー → Repositoryが呼ばれない",
			input: model.WordBook{
				ID:           1,
				Title:        "",
				Description:  "中学一年生の英単語",
				AdminId:      1,
				GenreId:      1,
				DifficultyId: 1,
				CreatedAt:    createdAt,
				UpdatedAt:    updatedAt,
			},
			mockRepoSetup: func(mockRepo *IWordBookRepositoryMock) {
				// バリデーションで失敗するため、Repositoryは呼ばれない
			},
			mockValidatorSetup: func(mockValidator *IWordBookValidatorMock) {
				mockValidator.On("ValidateTitle", mock.Anything).Return(errors.New("validate error"))
			},
			want: model.WordBookResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(IWordBookRepositoryMock)
			mockValidator := new(IWordBookValidatorMock)
			uc := NewWordBookUsecase(mockRepo, mockValidator)

			tt.mockRepoSetup(mockRepo)
			tt.mockValidatorSetup(mockValidator)

			got, err := uc.CreateWordBook(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("unexpected error: got %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unexpected result: got %+v, want %+v", got, tt.want)
			}

			mockRepo.AssertExpectations(t)
			mockValidator.AssertExpectations(t)
		})
	}
}
