package repository

import (
	"backend/model"

	"gorm.io/gorm"
)

type IAdminRepository interface {
	// 管理者を新規作成
	CreateAdmin(admin *model.Admin) error
	// メールアドレスで管理者を取得
	GetAdminByEmail(admin *model.Admin, email string) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) IAdminRepository {
	return &adminRepository{db}
}

// 管理者を新規作成
func (ar *adminRepository) CreateAdmin(admin *model.Admin) error {
	if err := ar.db.Create(admin).Error; err != nil {
		return err
	}
	return nil
}

// メールアドレスで管理者を取得
func (ar *adminRepository) GetAdminByEmail(admin *model.Admin, email string) error {
	if err := ar.db.Where("email = ?", email).First(admin).Error; err != nil {
		return err
	}
	return nil
}