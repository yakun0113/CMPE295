package main

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var SECRET_KEY = []byte("cmpe295")

func GenerateJWT(userEmail string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = userEmail
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}

func AuthRequired(c *gin.Context) {

}
