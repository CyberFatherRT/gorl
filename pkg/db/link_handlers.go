package db

import (
	"fmt"
	"log"
)

func GetLink(name string) (string, error) {
	var link string
	query := "SELECT link FROM url WHERE name = $1"

	result, err := DB.Query(query, name)

	if err != nil {
		return "", err
	}

	if !result.Next() {
		return "", fmt.Errorf("No link with that name")
	}

	DB.Exec("UPDATE url SET times = times + 1 WHERE name = $1", name)

	result.Scan(&link)
	return link, nil
}

func GenerateRandomLink(name, link string) (string, error) {
	query := "INSERT INTO url VALUES ($1, $2)"

	_, err := DB.Exec(query, name, link)

	if err != nil {
		return "", err
	}

	log.Printf("INFO: Create url %s -> %s", name, link)

	return name, nil
}
