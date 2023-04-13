package login

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/mariobac1/backend_webpages/model"
)

type Login struct {
	useCaseUser UseCaseUser
}

func New(uc UseCaseUser) Login {
	return Login{useCaseUser: uc}
}

// cambiamos la función para verificar las llaves
// func (l Login) Login(email, password, jwtSecretKey string) (model.User, string, error) {
func (l Login) Login(email, password string) (model.User, string, error) {
	user, err := l.useCaseUser.Login(email, password)
	if err != nil {
		return model.User{}, "", fmt.Errorf("%s %w", "useCaseUser.Login()", err)
	}

	claims := model.JWTCustomClaims{
		UserID: user.ID,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// data, err := token.SignedString([]byte(jwtSecretKey)) // acá en lugar de firmar con el []byte se firma con signKey de auth
	data, err := token.SignedString(signKey) // acá en lugar de firmar con el []byte se firma con signKey de auth
	if err != nil {
		return model.User{}, "", fmt.Errorf("%s %w", "token.SignedString()", err)
	}

	user.Password = ""

	return user, data, nil
}
