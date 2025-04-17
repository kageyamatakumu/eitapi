package usecase

import (
	"backend/model"
	"backend/repository"
	"backend/validator"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IAdminUsecase interface {
	CreateAdmin(admin model.Admin) (model.Admin, error)
	LoginAdmin(admin model.Admin) (string, error)
	UpdateAdminUserName(admin model.Admin, adminUserName string, userId uint) (model.AdminRes, error)
}

type adminUsecase struct {
	ar repository.IAdminRepository
	av validator.IAdminValidator
}

func NewAdminUsecase(ar repository.IAdminRepository, av validator.IAdminValidator) IAdminUsecase {
	return &adminUsecase{ar, av}
}

// 管理者を新規作成
func (au *adminUsecase) CreateAdmin(admin model.Admin) (model.Admin, error) {
	if err := au.av.ValidateEmail(admin.Email); err != nil {
		return model.Admin{}, err
	}

	if err := au.av.ValidatePassword(admin.Password); err != nil {
		return model.Admin{}, err
	}

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

// 管理者のログイン
func (au *adminUsecase) LoginAdmin(admin model.Admin) (string, error) {
	storedAdmin := model.Admin{}
	if err := au.ar.GetAdminByEmail(&storedAdmin, admin.Email); err != nil {
		return "", err
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedAdmin.Password), []byte(admin.Password))
	if err != nil {
		return "", errors.New("password does not match")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedAdmin.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 管理者のユーザー名を更新
func (au *adminUsecase) UpdateAdminUserName(admin model.Admin, adminUserName string, userId uint) (model.AdminRes, error) {
	if err := au.av.ValidateUserName(adminUserName); err != nil {
		return model.AdminRes{}, err
	}

	if err := au.ar.UpdateAdminUserName(&admin, adminUserName, userId); err != nil {
		return model.AdminRes{}, err
	}

	adminRes := model.AdminRes{
		ID:       admin.ID,
		Email:    admin.Email,
		UserName: admin.UserName,
	}

	return adminRes, nil
}
