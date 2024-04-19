package db

import (
	"database/sql"
	"fmt"
	util "gorl/pkg"
	"log"

	_ "github.com/lib/pq"
)

type DataBase struct {
	db      *sql.DB
	GenLink func(link string) (string, error)
	GetLink func(link string) (string, error)
}

var DB *DataBase

func Init() {
	pg_host := util.GetEnv("PG_HOST")
	pg_port := util.GetEnv("PG_PORT")
	pg_user := util.GetEnv("PG_USER")
	pg_password := util.GetEnv("PG_PASSWORD")
	pg_database := util.GetEnv("PG_DATABASE")

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pg_host, pg_port, pg_user, pg_password, pg_database)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatalf("ERROR: could not connect to database: %s", err)
	}

	DB = &DataBase{
		db:      db,
		GenLink: generateLink,
		GetLink: getLink,
	}

	err = DB.db.Ping()

	if err != nil {
		log.Fatalf("ERROR: could not connect to database: %s", err)
	}

	log.Printf("INFO: Connect to database at %s:%s", pg_host, pg_port)

	ensureRunned()
}

func ensureRunned() {
	_, err := DB.db.Exec(scheme)
	if err != nil {
		log.Fatal(err)
	}
}

func getLink(name string) (string, error) {
	query := `SELECT link FROM url WHERE name = $1`

	result, err := DB.db.Query(query, name)

	if err != nil {
		return "", err
	}

	if !result.Next() {
		return "", fmt.Errorf("Not link with that name")
	}

	var link string
	result.Scan(&link)

	return link, nil
}

func generateLink(link string) (string, error) {
	name := util.RandStringRunes(5)
	query := `INSERT INTO url VALUES ($1, $2)`

	_, err := DB.db.Exec(query, name, link)

	if err != nil {
		return "", err
	}

	log.Printf("INFO: Create url %s -> %s", name, link)

	return name, nil
}
