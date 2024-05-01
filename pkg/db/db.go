package db

import (
	"database/sql"
	"fmt"
	util "gorl/pkg"
	"log"

	_ "github.com/lib/pq"
)

type DataBase struct {
	db *sql.DB
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
		db: db,
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("ERROR: could not connect to database: %s", err)
	}

	ensureRunned()

	log.Printf("INFO: Connect to database at %s:%s", pg_host, pg_port)

}

func (db *DataBase) Exec(query string, args ...any) (sql.Result, error) {
	return db.db.Exec(query, args...)
}

func (db *DataBase) Query(query string, args ...any) (*sql.Rows, error) {
	return db.db.Query(query, args...)
}

func ensureRunned() {
	_, err := DB.Exec(scheme)
	if err != nil {
		log.Fatal(err)
	}
}
