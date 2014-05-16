package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Body string `json:"token"`
}

func NewToken(user *User) (*Token, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims["id"] = user.ID
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// TODO: Sign the token with an actual secret
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &Token{tokenString}, nil
}
