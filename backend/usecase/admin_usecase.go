package usecase

import (
	"backend/model"
	"backend/repository"

	"golang.org/x/crypto/bcrypt"
)

type IAdminUsecase interface {
	CreateAdmin(admin model.Admin) (model.Admin, error)
}

type adminUsecase struct {
	ar repository.IAdminRepository
}

func NewAdminUsecase(ar repository.IAdminRepository) IAdminUsecase {
	return &adminUsecase{ar}
}

// 管理者を新規作成
func (au *adminUsecase) CreateAdmin(admin model.Admin) (model.Admin, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return model.Admin{}, err
	}

	newUser := model.Admin{Email: admin.Email, Password: string(hash)}
	if err := au.ar.CreateAdmin(&newUser); err != nil {
		return model.Admin{}, err
	}

	return newUser, nil
}
