package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(dbUrl string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}
	log.Println("Connected to database")
	return db
}
