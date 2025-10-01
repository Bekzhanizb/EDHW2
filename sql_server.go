package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func main() {
	cfg := getDSNFromEnv()
	db, err := sql.Open("mysql", cfg)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
func getDSNFromEnv() string {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	pass := getEnv("DB_PASSWORD", "postgres")
	name := getEnv("DB_NAME", "assignment2_db")
	ssl := getEnv("DB_SSLMODE", "disable")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, name, ssl)
}
