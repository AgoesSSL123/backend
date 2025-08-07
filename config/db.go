package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal koneksi DB:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("DB tidak merespons:", err)
	}
	log.Println("✅ DB terkoneksi")
}
