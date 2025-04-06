package repository

import (
	"backend/model"
	"testing"
)

func Test_adminRepository_CreateAdmin(t *testing.T) {
	// テスト用DBを作成する
	db := SetupTestDB()
	// テスト後にDBをクリーンアップする
	defer func() {
		_ = db.Migrator().DropTable(&model.Admin{})
	}()

	repo := NewAdminRepository(db)

	tests := []struct {
		name    string
		admin   *model.Admin
		wantErr bool
	}{
		// TODO: Add test cases.
		{"success", &model.Admin{Email: "test@example.com", Password: "password123"}, false},
		{"duplicate email", &model.Admin{Email: "test@example.com", Password: "password123"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.CreateAdmin(tt.admin)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAdmin() error = %v, wantErr %v", err, tt.wantErr)
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

	tests := []struct {
		name    string
		admin   *model.Admin
		email   string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"success", &model.Admin{}, "test@example.com", false},
		{"not email", &model.Admin{}, "test2@example.com", true},
	}
	repo.CreateAdmin(&model.Admin{Email: "test@example.com", Password: "password123"})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.GetAdminByEmail(tt.admin, tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdminByEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
