package login

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/mariobac1/backend_webpages/model"
)

type Login struct {
	useCaseUser UseCaseUser
}

func New(uc UseCaseUser) Login {
	return Login{useCaseUser: uc}
}

func (l Login) Login(email, password, jwtSecretKey string) (model.User, string, error) {
	user, err := l.useCaseUser.Login(email, password)
	if err != nil {
		return model.User{}, "", fmt.Errorf("%s %w", "useCaseUser.Login()", err)
	}

	claims := model.JWTCustomClaims{
		UserID: user.ID,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	data, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return model.User{}, "", fmt.Errorf("%s %w", "token.SignedString()", err)
	}

	user.Password = ""

	return user, data, nil
}
