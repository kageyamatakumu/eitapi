package repository

import (
	"backend/model"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IAdminRepository interface {
	// 管理者を新規作成
	CreateAdmin(admin *model.Admin) error
	// メールアドレスで管理者を取得
	GetAdminByEmail(admin *model.Admin, email string) error
	// 管理者のユーザー名を更新
	UpdateAdminUserName(admin *model.Admin, adminUserName string, userId uint) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) IAdminRepository {
	return &adminRepository{db}
}

// 管理者を新規作成
func (ar *adminRepository) CreateAdmin(admin *model.Admin) error {
	if err := ar.db.Model(model.Admin{}).Create(admin).Error; err != nil {
		log.Printf("failed to create admin: %v\n", err)
		return err
	}

	log.Printf("successfully created admin\n")

	return nil
}

// メールアドレスで管理者を取得
func (ar *adminRepository) GetAdminByEmail(admin *model.Admin, email string) error {
	if err := ar.db.Model(model.Admin{}).Where("email = ?", email).First(admin).Error; err != nil {
		log.Printf("failed to get an admin by email %s: %v\n", email, err)
		return err
	}
	return nil
}

// 管理者のユーザー名を更新
func (ar *adminRepository) UpdateAdminUserName(admin *model.Admin, adminUserName string, userId uint) error {
	result := ar.db.Model(admin).Clauses(clause.Returning{}).Where("id = ?", userId).Update("user_name", adminUserName)

	if result.Error != nil {
		log.Printf("failed to update admin user name (userID: %d): %v\n", userId, result.Error)
		return result.Error
	}

	if result.RowsAffected < 1 {
		log.Printf("admin id %d not found\n", userId)
		return gorm.ErrRecordNotFound
	}

	log.Printf("successfully updated admin user name to '%s' (userID: %d)\n", admin.UserName, userId)

	return nil
}
