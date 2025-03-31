package usecase

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tryuuu/go-task-api/internal/domain/model"
	"github.com/tryuuu/go-task-api/internal/infrastructure"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key") // 本番では環境変数に！

func CreateUser(name, email, password string) (string, error) {
	// パスワードをハッシュ化
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := model.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}

	// DB保存
	result := infrastructure.DB.Create(&user)
	if result.Error != nil {
		return "", result.Error
	}

	// JWT生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
