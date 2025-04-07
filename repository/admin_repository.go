package repository

import (
	"backend/model"
	"log"

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
