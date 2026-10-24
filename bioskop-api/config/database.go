package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	// Cek environment variable dari Vercel
	connStr := os.Getenv("POSTGRES_URL")
	if connStr == "" {
		connStr = os.Getenv("DATABASE_URL")
	}

	// Jika variabel belum ada di Vercel, hindari koneksi ke localhost
	if connStr == "" {
		log.Println("WARNING: POSTGRES_URL atau DATABASE_URL tidak ditemukan di Environment Variables.")
		return nil
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
