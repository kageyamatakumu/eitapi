package usecase

import (
	"backend/model"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// --- モックの定義 ---

type IGenreRepositoryMock struct {
	mock.Mock
}

func (m *IGenreRepositoryMock) CreateGenre(genre *model.Genre) error {
	args := m.Called(genre)
	return args.Error(0)
}

func (m *IGenreRepositoryMock) UpdateGenre(genre *model.Genre, genreId uint) error {
	args := m.Called(genre, genreId)
	return args.Error(0)
}

func (m *IGenreRepositoryMock) DeleteGenre(genreId uint) error {
	args := m.Called(genreId)
	return args.Error(0)
}

// --- テスト ---

func Test_genreUsecase_CreateGenre(t *testing.T) {

	tests := []struct {
		name      string
		input     model.Genre
		mockSetup func(mockRepo *IGenreRepositoryMock)
		want      model.GenreRes
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name:  "正常系: ジャンル作成成功",
			input: model.Genre{GenreName: "日常会話"},
			mockSetup: func(mockRepo *IGenreRepositoryMock) {
				mockRepo.On("CreateGenre", mock.Anything).Run(func(args mock.Arguments) {
					ptr := args.Get(0).(*model.Genre)
					ptr.ID = 1
				}).Return(nil)
			},
			want:    model.GenreRes{ID: 1, GenreName: "日常会話"},
			wantErr: false,
		},
		{
			name:  "異常系: ジャンル作成失敗",
			input: model.Genre{GenreName: "日常会話"},
			mockSetup: func(mockRepo *IGenreRepositoryMock) {
				mockRepo.On("CreateGenre", mock.Anything).Return(errors.New("db error"))
			},
			want:    model.GenreRes{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(IGenreRepositoryMock)

			tt.mockSetup(mockRepo)

			uc := NewGenreUsecase(mockRepo)

			got, err := uc.CreateGenre(tt.input)

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

func Test_genreUsecase_UpdateGenre(t *testing.T) {
	tests := []struct {
		name      string
		input     model.Genre
		inputId   uint
		mockSetup func(mockRepo *IGenreRepositoryMock)
		want      model.GenreRes
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name:    "正常系: ジャンル更新成功",
			input:   model.Genre{GenreName: "ビジネス会話"},
			inputId: 1,
			mockSetup: func(mockRepo *IGenreRepositoryMock) {
				mockRepo.On("UpdateGenre", mock.Anything, uint(1)).Run(func(args mock.Arguments) {
					ptr := args.Get(0).(*model.Genre)
					*ptr = model.Genre{ID: 1, GenreName: "ビジネス会話"}
				}).Return(nil)
			},
			want:    model.GenreRes{ID: 1, GenreName: "ビジネス会話"},
			wantErr: false,
		},
		{
			name:    "異常系: ジャンル更新失敗",
			input:   model.Genre{GenreName: "ビジネス会話"},
			inputId: 1,
			mockSetup: func(mockRepo *IGenreRepositoryMock) {
				mockRepo.On("UpdateGenre", mock.Anything, uint(1)).Return(errors.New("db error"))
			},
			want:    model.GenreRes{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(IGenreRepositoryMock)

			tt.mockSetup(mockRepo)

			uc := NewGenreUsecase(mockRepo)

			got, err := uc.UpdateGenre(tt.input, tt.inputId)

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

func Test_genreUsecase_DeleteGenre(t *testing.T) {
	tests := []struct {
		name      string
		inputId   uint
		mockSetup func(mockRepo *IGenreRepositoryMock)
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name:    "正常系: ジャンル削除成功",
			inputId: 1,
			mockSetup: func(mockRepo *IGenreRepositoryMock) {
				mockRepo.On("DeleteGenre", uint(1)).Return(nil)
			},
			wantErr: false,
		},
		{
			name:    "異常系: DBエラーが発生",
			inputId: 1,
			mockSetup: func(mockRepo *IGenreRepositoryMock) {
				mockRepo.On("DeleteGenre", uint(1)).Return(errors.New("db error")).Once()
			},
			wantErr: true,
		},
		{
			name:    "異常系: 削除対象が存在せず RecordNotFound エラー",
			inputId: 999,
			mockSetup: func(mockRepo *IGenreRepositoryMock) {
				mockRepo.On("DeleteGenre", uint(999)).Return(gorm.ErrRecordNotFound)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(IGenreRepositoryMock)

			tt.mockSetup(mockRepo)

			uc := NewGenreUsecase(mockRepo)

			err := uc.DeleteGenre(tt.inputId)

			if (err != nil) != tt.wantErr {
				t.Errorf("unexpected error: err %v, wantErr %v", err, tt.wantErr)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
