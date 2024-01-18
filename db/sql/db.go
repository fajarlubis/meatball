package sql

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	dbSource = "postgresql://root:meatball@localhost:1302/meatball?sslmode=disable"
)

type DB struct {
	pool *sql.DB
}

var db = &DB{}

func NewConnection() *DB {
	postgres, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal(err)
	}

	if err := postgres.Ping(); err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	postgres.SetConnMaxIdleTime(time.Minute)
	postgres.SetConnMaxLifetime(time.Minute * 5)

	db.pool = postgres

	return db
}
