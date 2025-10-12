package usecase

import (
	"strings"
	"time"
	"wheels-api/config"
	"wheels-api/model"
	"wheels-api/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Claims representa as reivindicações do JWT.
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

type UserUseCase struct {
	repo repository.UserRepository
}

func (uc *UserUseCase) Register(user model.User) (int64, error) {
	user.Email = strings.ToLower(user.Email)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user.Password = string(hashedPassword)
	return uc.repo.CreateUser(user)
}

func (uc *UserUseCase) Login(email, password string) (string, error) {
	email = strings.ToLower(email)
	user, err := uc.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.GetJWTSecret())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
