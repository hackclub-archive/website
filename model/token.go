package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Token represents a JSON Web Token (JWT) as returned to the user.
type Token struct {
	Body string `json:"token"`
}

// NewToken creates a new Token from a provided user.
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
