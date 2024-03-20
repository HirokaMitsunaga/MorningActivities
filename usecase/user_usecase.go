package usecase

import (
	"go-api/model"
	"go-api/repository"
	"go-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	//返り値のstringはJWTトークン
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	userRes := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return userRes, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}
	//id（Email）が存在しており、入力されたパスワードが存在していればOK
	storedUser := model.User{}
	//データベースからメールアドレスのユーザーがいるのか確認して、storedUserへ格納する
	//ユーザがいなければエラーを出す。
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	//平文で送られてきたパスワードとデータベース内のパスワード（復号かしたもの）を比較する
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	//JWTトークンの設定
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	//JWTトークンの生成
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, err
}
