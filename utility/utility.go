package utility

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("Scanning-SecretKey")

func (r *TokenReq) CreateJwtToken() (*TokenReq, error) {

	claims := jwt.MapClaims{
		"id":  r.Id,
		"exp": time.Now().Add(time.Hour * 5).Unix(),
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error creating JWT:", err)
		return nil, err
	}

	r.Token = tokenString

	return r, nil

}
