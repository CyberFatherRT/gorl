package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	util "gorl/pkg"
	"log"
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

	err = DB.db.Ping()

	if err != nil {
		log.Fatalf("ERROR: could not connect to database: %s", err)
	}

	log.Printf("INFO: Connect to database at %s:%s", pg_host, pg_port)
}

func GetLink() string {
	return util.RandStringRunes(5)
}
