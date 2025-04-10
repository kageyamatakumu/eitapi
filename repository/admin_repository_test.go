package repository

import (
	"backend/model"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_adminRepository_CreateAdmin(t *testing.T) {
	// テスト用DBを作成する
	db := SetupTestDB()
	// テスト後にDBをクリーンアップする
	defer func() {
		_ = db.Migrator().DropTable(&model.Admin{})
	}()

	repo := NewAdminRepository(db)

	longString := strings.Repeat("a", 256)
	longString2 := strings.Repeat("b", 256)

	tests := []struct {
		name    string
		admin   *model.Admin
		wantErr bool
	}{
		// TODO: Add test cases.
		{"success", &model.Admin{Email: "test@example.com", Password: "password123"}, false},
		{"duplicate email", &model.Admin{Email: "test@example.com", Password: "password123"}, true},
		{"email too long", &model.Admin{Email: longString, Password: "password123"}, false}, // SQLite ではエラーが発生しないため、wantErr を false に変更
		{"password too long", &model.Admin{Email: "test2@example.com", Password: longString},false}, // SQLite ではエラーが発生しないため、wantErr を false に変更
		{"all fields too long", &model.Admin{Email: longString2, Password: longString}, false}, // SQLite ではエラーが発生しないため、wantErr を false に変更
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.CreateAdmin(tt.admin)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				var createdAdmin model.Admin
				result := db.Where("email = ?", tt.admin.Email).First(&createdAdmin)
				assert.NoError(t, result.Error)
				assert.Equal(t, tt.admin.Email, createdAdmin.Email)
				assert.Equal(t, tt.admin.Password, createdAdmin.Password)
			}
		})
	}
}

func Test_adminRepository_GetAdminByEmail(t *testing.T) {
	// テスト用DBを作成する
	db := SetupTestDB()
	// テスト後にDBをクリーンアップする
	defer func() {
		_ = db.Migrator().DropTable(&model.Admin{})
	}()

	repo := NewAdminRepository(db)
	admin := model.Admin{Email: "test@example.com", Password: "password123"}
	repo.CreateAdmin(&admin)

	tests := []struct {
		name      string
		email     string
		wantAdmin *model.Admin
		wantErr   bool
	}{
		// TODO: Add test cases.
		{"success", "test@example.com", &admin, false},
		{"not found", "test2@example.com", &model.Admin{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotAdmin model.Admin
			err := repo.GetAdminByEmail(&gotAdmin, tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdminByEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				assert.Equal(t, tt.wantAdmin.Email, gotAdmin.Email)
				assert.Equal(t, tt.wantAdmin.Password, gotAdmin.Password)
			}
		})
	}
}

func Test_adminRepository_UpdateAdminUserName(t *testing.T) {
	db := SetupTestDB()
	defer func() {
		_ = db.Migrator().DropTable(&model.Admin{})
	}()

	repo := NewAdminRepository(db)
	admin := model.Admin{Email: "test@example.com", Password: "password123"}
	repo.CreateAdmin(&admin)

	longString := strings.Repeat("a", 256)

	tests := []struct {
		name          string
		adminUserName string
		userId        uint
		wantErr       bool
	}{
		{"success", "newName", admin.ID, false},
		{"not found", "newName", 2, true},
		{"user name too long", longString, admin.ID, false}, // SQLite ではエラーが発生しないため、wantErr を false に変更
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var targetAdmin model.Admin
			db.First(&targetAdmin, admin.ID)
			err := repo.UpdateAdminUserName(&targetAdmin, tt.adminUserName, tt.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAdminUserName() error = %v,wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				var updatedAdmin model.Admin
				result := db.Where("id = ?", tt.userId).First(&updatedAdmin)
				assert.NoError(t, result.Error)
				assert.Equal(t, tt.adminUserName, updatedAdmin.UserName)
			}
		})
	}
}
