package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	// Ambil string koneksi dari Vercel
	connStr := os.Getenv("POSTGRES_URL")
	if connStr == "" {
		connStr = os.Getenv("DATABASE_URL")
	}

	// JIKA KOSONG: Langsung stop di sini, JANGAN coba ke 127.0.0.1
	if connStr == "" {
		log.Println("[INFO] POSTGRES_URL / DATABASE_URL belum terbaca dari Vercel Env.")
		return nil
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("[ERROR] sql.Open gagal:", err)
		return nil
	}

	if err = db.Ping(); err != nil {
		log.Println("[ERROR] db.Ping gagal:", err)
		return nil
	}

	log.Println("[SUCCESS] Terhubung ke PostgreSQL Neon!")
	return db
}
