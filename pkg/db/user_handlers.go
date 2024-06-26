package db

import (
	"github.com/matthewhartstonge/argon2"
	"log"
)

var Argon2 argon2.Config

func CreateUser(username, password string, isAdmin bool) error {
	query := "INSERT INTO users(username, password, is_admin) VALUES ($1, $2, $3)"
	hashedPassword, _ := Argon2.HashEncoded([]byte(password))

	_, err := DB.Query(query, username, hashedPassword, isAdmin)
	if err != nil {
		return err
	}

	log.Printf("INFO: Create user `%s`", username)

	return nil
}

func UpdatePassword(id int, password string) error {
	query := "UPDATE users SET password = $1 WHERE id = $2"

	_, err := DB.Query(query, password, id)
	if err != nil {
		return err
	}

	log.Printf("INFO: Update `password` of user `%d`", id)

	return nil
}

func UpdateAdminUser(id int, isAdmin bool) error {
	query := "UPDATE users SET is_admin = $1 WHERE id = $2"

	_, err := DB.Query(query, isAdmin, id)
	if err != nil {
		return err
	}

	log.Printf("INFO: Update `isAdmin` field of user `%d`. isAdmin = %t", id, isAdmin)

	return nil
}
