package pkg

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var JWTsecret []byte

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func SignedJwt(username string, isAdmin bool) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).UnixMilli()
	claims["authorized"] = true
	claims["username"] = username
	claims["isAdmin"] = isAdmin

	tokenString, err := token.SignedString(JWTsecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetEnv(key string) string {
	variable := os.Getenv(key)
	if len(variable) == 0 {
		log.Fatalf("ERROR: %s envirnment variable not set.", key)
	}
	return variable
}
