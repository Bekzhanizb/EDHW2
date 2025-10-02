package main

import (
	"database/sql"
	"log"
)

func main() {
	connect := getDSNFromEnv()
	db, err := sql.Open("postgres", connect)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	if err := InitSQLSchema(db); err != nil {
		log.Fatalf("Error initializing scheme: %v", err)
	}
}
func InitSQLSchema(db *sql.DB) error {
	q := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		age INTEGER NOT NULL
	);
	`
	_, err := db.Exec(q)
	return err
}
