package usecase

import (
	"backend/model"
	"backend/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IAdminUsecase interface {
	CreateAdmin(admin model.Admin) (model.Admin, error)
	LoginAdmin(admin model.Admin) (string, error)
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

// 管理者のログイン
func (au *adminUsecase) LoginAdmin(admin model.Admin) (string, error) {
	storedAdmin := model.Admin{}
	if err := au.ar.GetAdminByEmail(&storedAdmin, admin.Email); err != nil {
		return "", err
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedAdmin.Password), []byte(admin.Password))
	if err != nil {
		return "", err
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