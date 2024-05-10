package utils

import (
	"ERP-API/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("somethingstrangehappens")

type JWTClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(enterprise models.Enterprise) (string, error) {
	claims := jwt.MapClaims{}
	claims["enterprise_id"] = enterprise.ID
	claims["enterprise_name"] = enterprise.Name
	claims["enterprise_email"] = enterprise.Email
	claims["enterprise_subscription_name"] = enterprise.Sub

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
