package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secretKey = []byte("somethingstrangehappens")

type JWTClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(enterpriseId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["enterprise_id"] = enterpriseId
	claims["expiration_period"] = time.Now().Add(time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return err
	}

	return nil
}
