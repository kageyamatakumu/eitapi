package usecase

import (
	"backend/model"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

// --- モックの定義 ---

type IDifficultyRepositoryMock struct {
	mock.Mock
}

func (m *IDifficultyRepositoryMock) GetAllDifficulties(difficulties *[]model.Difficulty) error {
	args := m.Called(difficulties)
	return args.Error(0)

	// mockRepo := new(IDifficultyRepositoryMock)
	// mockRepo.On("GetAllDifficulties", mock.Anything).Run(func(args mock.Arguments) {
	// 	ptr := args.Get(0).(*[]model.Difficulty)
	// 	*ptr = []model.Difficulty{
	// 		{ID: 1, DifficultyLevel: "Easy"},
	// 		{ID: 2, DifficultyLevel: "Hard"},
	// 	}
	// }).Return(nil)
}

func (m *IDifficultyRepositoryMock) CreateDifficulty(difficulty *model.Difficulty) error {
	args := m.Called(difficulty)
	return args.Error(0)

	// mockRepo := new(IDifficultyRepositoryMock)
	// mockRepo.On("CreateDifficulty", mock.Anything).Return(nil)
}

func (m *IDifficultyRepositoryMock) UpdateDifficulty(difficulty *model.Difficulty, difficultyId uint) error {
	args := m.Called(difficulty, difficultyId)
	return args.Error(0)

	// mockRepo := new(IDifficultyRepositoryMock)
	// mockRepo.On("UpdateDifficulty",  mock.Anything, mock.Anything).Return(nil)
}

func (m *IDifficultyRepositoryMock) DeleteDifficulty(difficultyId uint) error {
	args := m.Called(difficultyId)
	return args.Error(0)

	// mockRepo := new(IDifficultyRepositoryMock)
	// mockRepo.On("DeleteDifficulty", mock.Anything).Return(nil)
	// mockRepo.On("DeleteDifficulty", uint(1)).Return(nil)
	// mockRepo.On("DeleteDifficulty", uint(999)).Return(errors.New("not found"))
}

type IDifficultyValidatorMock struct {
	mock.Mock
}

func (m *IDifficultyValidatorMock) ValidateDifficultyLevel(difficultyLevel model.DifficultyLevel) error {
	args := m.Called(difficultyLevel)
	return args.Error(0)
}

// --- テスト ---

func Test_difficultyUsecase_GetAllDifficulties(t *testing.T) {

	tests := []struct {
		name      string
		mockSetup func(mockRepo *IDifficultyRepositoryMock)
		want      []model.DifficultyResponse
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name: "正常系: 難易度リスト取得成功",
			mockSetup: func(mockRepo *IDifficultyRepositoryMock) {
				mockRepo.On("GetAllDifficulties", mock.Anything).Run(func(args mock.Arguments) {
					ptr := args.Get(0).(*[]model.Difficulty)
					*ptr = []model.Difficulty{
						{ID: 1, DifficultyLevel: model.Easy},
						{ID: 2, DifficultyLevel: model.Hard},
					}
				}).Return(nil)
			},
			want: []model.DifficultyResponse{
				{ID: 1, DifficultyLevel: model.Easy},
				{ID: 2, DifficultyLevel: model.Hard},
			},
			wantErr: false,
		},
		{
			name: "異常系: リポジトリエラー発生",
			mockSetup: func(mockRepo *IDifficultyRepositoryMock) {
				mockRepo.On("GetAllDifficulties", mock.Anything).Return(errors.New("db error"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(IDifficultyRepositoryMock)
			mockValidator := new(IDifficultyValidatorMock)
			uc := NewDifficultyUsecase(mockRepo, mockValidator)

			tt.mockSetup(mockRepo)

			got, err := uc.GetAllDifficulties()

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

func Test_difficultyUsecase_CreateDifficulty(t *testing.T) {
	tests := []struct {
		name               string
		input              model.Difficulty
		mockRepoSetup      func(mockRepo *IDifficultyRepositoryMock)
		mockValidatorSetup func(mockValidator *IDifficultyValidatorMock)
		want               model.DifficultyResponse
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			name:  "正常系: 難易度作成成功",
			input: model.Difficulty{ID: 1, DifficultyLevel: model.Easy},
			mockRepoSetup: func(mockRepo *IDifficultyRepositoryMock) {
				mockRepo.On("CreateDifficulty", mock.Anything).Return(nil)
			},
			mockValidatorSetup: func(mockValidator *IDifficultyValidatorMock) {
				mockValidator.On("ValidateDifficultyLevel", mock.Anything).Return(nil)
			},
			want:    model.DifficultyResponse{ID: 1, DifficultyLevel: model.Easy},
			wantErr: false,
		},
		{
			name:  "異常系: 難易度作成失敗",
			input: model.Difficulty{ID: 1, DifficultyLevel: model.Easy},
			mockRepoSetup: func(mockRepo *IDifficultyRepositoryMock) {
				mockRepo.On("CreateDifficulty", mock.Anything).Return(errors.New("db error"))
			},
			mockValidatorSetup: func(mockValidator *IDifficultyValidatorMock) {
				mockValidator.On("ValidateDifficultyLevel", mock.Anything).Return(nil)
			},
			want:    model.DifficultyResponse{},
			wantErr: true,
		},
		{
			name:  "異常系: 不正なDifficultyLevelでバリデーションエラー → Repositoryが呼ばれない",
			input: model.Difficulty{ID: 2, DifficultyLevel: 4},
			mockRepoSetup: func(mockRepo *IDifficultyRepositoryMock) {
				// バリデーションで失敗するため、Repositoryは呼ばれない
			},
			mockValidatorSetup: func(mockValidator *IDifficultyValidatorMock) {
				mockValidator.On("ValidateDifficultyLevel", mock.Anything).Return(errors.New("validate error"))
			},
			want:    model.DifficultyResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(IDifficultyRepositoryMock)
			mockValidator := new(IDifficultyValidatorMock)

			tt.mockRepoSetup(mockRepo)
			tt.mockValidatorSetup(mockValidator)

			uc := NewDifficultyUsecase(mockRepo, mockValidator)

			got, err := uc.CreateDifficulty(tt.input)

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

func Test_difficultyUsecase_UpdateDifficulty(t *testing.T) {
	tests := []struct {
		name               string
		input              model.Difficulty
		inputId            uint
		mockRepoSetup      func(mockRepo *IDifficultyRepositoryMock)
		mockValidatorSetup func(mockValidator *IDifficultyValidatorMock)
		want               model.DifficultyResponse
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			name:         "正常系: 難易度更新成功",
			input:        model.Difficulty{DifficultyLevel: model.Hard},
			inputId:      1,
			mockRepoSetup: func(mockRepo *IDifficultyRepositoryMock) {
				mockRepo.On("UpdateDifficulty", mock.Anything, uint(1)).Run(func(args mock.Arguments) {
					ptr := args.Get(0).(*model.Difficulty)
					*ptr = model.Difficulty{ID: 1, DifficultyLevel: model.Hard}
				}).Return(nil)
			},
			mockValidatorSetup: func(mockValidator *IDifficultyValidatorMock) {
				mockValidator.On("ValidateDifficultyLevel", mock.Anything).Return(nil)
			},
			want:    model.DifficultyResponse{ID: 1, DifficultyLevel: model.Hard},
			wantErr: false,
		},
		{
			name:    "異常系: DBエラーが発生",
			input:   model.Difficulty{DifficultyLevel: model.Hard},
			inputId: 99,
			mockRepoSetup: func(mockRepo *IDifficultyRepositoryMock) {
				mockRepo.On("UpdateDifficulty", mock.Anything, uint(99)).Return(errors.New("db error"))
			},
			mockValidatorSetup: func(mockValidator *IDifficultyValidatorMock) {
				mockValidator.On("ValidateDifficultyLevel", model.Hard).Return(nil)
			},
			want:    model.DifficultyResponse{},
			wantErr: true,
		},
		{
			name:    "異常系: 不正なDifficultyLevelでバリデーションエラー → Repositoryが呼ばれない",
			input:   model.Difficulty{DifficultyLevel: model.Hard},
			inputId: 1,
			mockRepoSetup: func(mockRepo *IDifficultyRepositoryMock) {
				// バリデーションで失敗するため、Repositoryは呼ばれない
			},
			mockValidatorSetup: func(mockValidator *IDifficultyValidatorMock) {
				mockValidator.On("ValidateDifficultyLevel", model.Hard).Return(errors.New("validate error"))
			},
			want:    model.DifficultyResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(IDifficultyRepositoryMock)
			mockValidator := new(IDifficultyValidatorMock)

			tt.mockRepoSetup(mockRepo)
			tt.mockValidatorSetup(mockValidator)

			uc := NewDifficultyUsecase(mockRepo, mockValidator)

			got, err := uc.UpdateDifficulty(tt.input, tt.inputId)

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
