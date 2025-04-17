package usecase

import (
	"backend/model"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// --- モックの定義 ---

type IAdminRepositoryMock struct {
	mock.Mock
}

func (m *IAdminRepositoryMock) CreateAdmin(admin *model.Admin) error {
	args := m.Called(admin)
	return args.Error(0)
}

func (m *IAdminRepositoryMock) GetAdminByEmail(admin *model.Admin, email string) error {
	args := m.Called(admin, email)
	return args.Error(0)
}

func (m *IAdminRepositoryMock) UpdateAdminUserName(admin *model.Admin, adminUserName string, userId uint) error {
	args := m.Called(admin, adminUserName, userId)
	return args.Error(0)
}

type IAdminValidatorMock struct {
	mock.Mock
}

func (m *IAdminValidatorMock) ValidateEmail(email string) error {
	args := m.Called(email)
	return args.Error(0)
}

func (m *IAdminValidatorMock) ValidatePassword(password string) error {
	args := m.Called(password)
	return args.Error(0)
}

func (m *IAdminValidatorMock) ValidateUserName(userName string) error {
	args := m.Called(userName)
	return args.Error(0)
}

// --- テスト ---
var rawPassword string = "securePass123"

func Test_adminUsecase_CreateAdmin(t *testing.T) {
	mockRepo := new(IAdminRepositoryMock)
	mockValidator := new(IAdminValidatorMock)

	// Usecaseのインスタンスを作成
	uc := NewAdminUsecase(mockRepo, mockValidator)

	tests := []struct {
		name          string
		admin         model.Admin
		validateEmail error
		validatePass  error
		createFunc    func(admin *model.Admin) error
		wantErr       bool
		wantErrMsg    string
	}{
		{
			name: "Successfully create admin (valid email and password)",
			admin: model.Admin{
				Email:    "test@example.com",
				Password: rawPassword,
			},
			validateEmail: nil,
			validatePass:  nil,
			createFunc: func(admin *model.Admin) error {
				// モックの処理内でIDをセット
				admin.ID = 1
				return nil
			},
			wantErr:    false,
			wantErrMsg: "",
		},
		{
			name: "Email validation error",
			admin: model.Admin{
				Email:    "invalid-email",
				Password: "pass",
			},
			validateEmail: errors.New("invalid email"),
			validatePass:  nil,
			createFunc:    nil,
			wantErr:       true,
			wantErrMsg:    "invalid email",
		},
		{
			name: "Password validation error",
			admin: model.Admin{
				Email:    "pass",
				Password: "invalid-password",
			},
			validateEmail: nil,
			validatePass:  errors.New("invalid password"),
			createFunc:    nil,
			wantErr:       true,
			wantErrMsg:    "invalid password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// バリデーションをモック
			mockValidator.On("ValidateEmail", tt.admin.Email).Return(tt.validateEmail)

			if tt.name == "Email validation error" {
				// ValidatePassword のモックは設定しない
			} else {
				mockValidator.On("ValidatePassword", tt.admin.Password).Return(tt.validatePass)
			}

			// CreateAdmin メソッドのモックを設定
			if tt.createFunc != nil {
				mockRepo.On("CreateAdmin", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
					// ここでテストケースの処理を実行
					tt.createFunc(args.Get(0).(*model.Admin))
				})
			} else {
				// createFunc が設定されていない場合、単に nil を返す
				mockRepo.On("CreateAdmin", mock.Anything).Return(nil)
			}

			// テスト実行
			adminCopy := tt.admin
			got, err := uc.CreateAdmin(adminCopy)

			// エラーが発生すべきか確認
			if tt.wantErr {
				assert.Error(t, err)
				if tt.wantErrMsg != "" {
					assert.EqualError(t, err, tt.wantErrMsg)
				}
			} else {
				// エラーが発生しなかった場合
				assert.NoError(t, err)
				// メールアドレスが一致しているか
				assert.Equal(t, tt.admin.Email, got.Email)
				// パスワードがハッシュ化されているか
				assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(got.Password), []byte(rawPassword)))
				// IDがセットされたか
				assert.Equal(t, uint(1), got.ID)
			}

			// モックが正しく呼び出されたか確認
			mockRepo.AssertExpectations(t)
			mockValidator.AssertExpectations(t)
		})
	}
}

func TestAdminUsecase_LoginAdmin(t *testing.T) {
	mockRepo := new(IAdminRepositoryMock)
	mockValidator := new(IAdminValidatorMock)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)

	uc := NewAdminUsecase(mockRepo, mockValidator)

	// 環境変数 SECRET を設定
	os.Setenv("SECRET", "test-secret")
	defer os.Unsetenv("SECRET")

	tests := []struct {
		name          string
		admin         model.Admin
		storedAdmin   model.Admin
		getAdminError error
		wantErr       bool
		wantErrMsg    string
	}{
		{
			name: "Successfully login admin",
			admin: model.Admin{
				Email:    "test@example.com",
				Password: rawPassword,
			},
			storedAdmin: model.Admin{
				ID:       1,
				Email:    "test@example.com",
				Password: string(hashedPassword),
			},
			getAdminError: nil,
			wantErr:       false,
			wantErrMsg:    "",
		},
		{
			name: "Admin not found",
			admin: model.Admin{
				Email:    "nonexistent@example.com",
				Password: rawPassword,
			},
			storedAdmin:   model.Admin{},
			getAdminError: errors.New("admin not found"),
			wantErr:       true,
			wantErrMsg:    "admin not found",
		},
		{
			name: "Incorrect password",
			admin: model.Admin{
				Email:    "test@example.com",
				Password: "wrong-password",
			},
			storedAdmin: model.Admin{
				ID:       1,
				Email:    "test@example.com",
				Password: string(hashedPassword),
			},
			getAdminError: nil,
			wantErr:       true,
			wantErrMsg:    "password does not match",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.getAdminError != nil {
				// Admin not found の場合
				mockRepo.On("GetAdminByEmail", mock.AnythingOfType("*model.Admin"), tt.admin.Email).
					Return(tt.getAdminError)
			} else {
				// 成功 or パスワード不一致の場合
				mockRepo.On("GetAdminByEmail", mock.AnythingOfType("*model.Admin"), tt.admin.Email).
					Return(nil).Run(func(args mock.Arguments) {
					adminPtr := args.Get(0).(*model.Admin)
					*adminPtr = tt.storedAdmin
				})
			}

			adminCopy := tt.admin
			got, err := uc.LoginAdmin(adminCopy)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErrMsg)
				assert.Empty(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, got)

				token, err := jwt.Parse(got, func(token *jwt.Token) (interface{}, error) {
					return []byte(os.Getenv("SECRET")), nil
				})

				assert.NoError(t, err)
				claims, ok := token.Claims.(jwt.MapClaims)
				assert.True(t, ok)
				assert.Equal(t, float64(tt.storedAdmin.ID), claims["user_id"])
				exp := int64(claims["exp"].(float64))
				assert.Greater(t, exp, time.Now().Unix())
			}

			mockRepo.AssertExpectations(t)
		})
	}

}

func Test_adminUsecase_UpdateAdminUserName(t *testing.T) {
	mockRepo := new(IAdminRepositoryMock)
	mockValidator := new(IAdminValidatorMock)

	uc := NewAdminUsecase(mockRepo, mockValidator)

	tests := []struct {
		name          string
		admin         model.Admin
		adminUserName string
		userId        uint
		want          model.AdminRes
		validateName  error
		wantErr       bool
		wantErrMsg    string
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			admin: model.Admin{
				ID:       1,
				Email:    "test@example.com",
				UserName: "oldName",
			},
			adminUserName: "newName",
			userId:        1,
			want: model.AdminRes{
				ID:       1,
				Email:    "test@example.com",
				UserName: "newName",
			},
			validateName: nil,
			wantErr:      false,
			wantErrMsg:   "",
		},
		{
			name: "Validation error",
			admin: model.Admin{
				ID:       1,
				Email:    "test@example.com",
				UserName: "oldName",
			},
			adminUserName: "badName!",
			userId:        1,
			want:          model.AdminRes{},
			validateName:  errors.New("invalid user name"),
			wantErr:       true,
			wantErrMsg:    "invalid user name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockValidator.On("ValidateUserName", tt.adminUserName).Return(tt.validateName)

			// バリデーションエラーがなければ、Repoのモックも設定
			if tt.validateName == nil {
				mockRepo.On("UpdateAdminUserName", mock.AnythingOfType("*model.Admin"), tt.adminUserName, tt.userId).
					Run(func(args mock.Arguments) {
						argAdmin := args.Get(0).(*model.Admin)
						argAdmin.UserName = tt.adminUserName
					}).
					Return(nil)
			}

			adminCopy := tt.admin
			got, err := uc.UpdateAdminUserName(adminCopy, tt.adminUserName, tt.userId)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErrMsg)
				assert.Equal(t, model.AdminRes{}, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}

			mockValidator.AssertExpectations(t)
			mockRepo.AssertExpectations(t)
		})
	}
}
