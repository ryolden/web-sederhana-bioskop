package repository

import (
	"bioskop-api/entity"
	"context"
	"database/sql"
)

type BioskopRepository interface {
	Create(ctx context.Context, bioskop *entity.Bioskop) error
}

type bioskopRepository struct {
	db *sql.DB
}

func NewBioskopRepository(db *sql.DB) BioskopRepository {
	return &bioskopRepository{db: db}
}

func (r *bioskopRepository) Create(ctx context.Context, bioskop *entity.Bioskop) error {
	query := `INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, bioskop.Nama, bioskop.Lokasi, bioskop.Rating).Scan(&bioskop.ID)
	if err != nil {
		return err
	}
	return nil
}
