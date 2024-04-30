package pkg

import (
	"log"
	"math/rand"
	"os"
)

var letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetEnv(key string) string {
	variable := os.Getenv(key)
	if len(variable) == 0 {
		log.Fatalf("ERROR: %s envirnment variable not set.", key)
	}
	return variable
}
