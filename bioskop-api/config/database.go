package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	// 1. Cek beberapa kemungkinan nama variabel koneksi Vercel Postgres
	connStr := os.Getenv("POSTGRES_URL")
	if connStr == "" {
		connStr = os.Getenv("DATABASE_URL")
	}
	if connStr == "" {
		connStr = os.Getenv("POSTGRES_URL_NON_POOLING")
	}

	// 2. Jika semua variabel Vercel kosong, baru gunakan fallback local
	if connStr == "" {
		log.Println("WARNING: Environment variable database tidak ditemukan. Menggunakan fallback local.")
		host := getEnv("DB_HOST", "localhost")
		port := getEnv("DB_PORT", "5432")
		user := getEnv("DB_USER", "postgres")
		password := getEnv("DB_PASSWORD", "postgres")
		dbname := getEnv("DB_NAME", "bioskop_db")

		connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Gagal sql.Open:", err)
		return nil
	}

	if err = db.Ping(); err != nil {
		log.Println("Gagal db.Ping ke database:", err)
		return nil
	}

	log.Println("Berhasil terhubung ke database!")
	return db
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
