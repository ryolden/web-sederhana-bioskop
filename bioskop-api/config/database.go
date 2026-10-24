package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connStr := os.Getenv("POSTGRES_URL")
	if connStr == "" {
		connStr = os.Getenv("DATABASE_URL")
	}

	// Jika kosong, stop tanpa menyentuh 127.0.0.1
	if connStr == "" {
		log.Println("[WARN] POSTGRES_URL / DATABASE_URL belum terbaca di Vercel.")
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

	log.Println("[SUCCESS] Terhubung ke Postgres Neon!")
	return db
}
