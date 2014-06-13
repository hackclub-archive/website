package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Token represents a JSON Web Token (JWT) as returned to the user.
type Token struct {
	UserID   int64     `json:"id"`
	UserType string    `json:"type"`
	Token    string    `json:"token"`
	Expires  time.Time `json:"expires"`
}

// NewToken creates a new Token from a provided user.
func NewToken(user *User) (*Token, error) {
	expires := time.Now().Add(time.Hour * 72)

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims["id"] = user.ID
	token.Claims["exp"] = expires.Unix()

	// TODO: Sign the token with an actual secret
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &Token{UserID: user.ID, UserType: user.Type, Token: tokenString,
		Expires: expires}, nil
}
