package config

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connStr := os.Getenv("POSTGRES_URL")
	if connStr == "" {
		connStr = os.Getenv("DATABASE_URL")
	}

	if connStr == "" {
		log.Println("[WARN] Variable POSTGRES_URL / DATABASE_URL belum diatur.")
		return nil
	}

	// Hapus channel_binding otomatis agar kompatibel dengan driver Go lib/pq
	connStr = strings.ReplaceAll(connStr, "channel_binding=require&", "")
	connStr = strings.ReplaceAll(connStr, "&channel_binding=require", "")
	connStr = strings.ReplaceAll(connStr, "?channel_binding=require", "?")

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
