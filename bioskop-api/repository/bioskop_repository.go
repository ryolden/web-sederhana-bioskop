package repository

import (
	"bioskop-api/entity"
	"context"
	"database/sql"
	"errors"
)

type BioskopRepository interface {
	Create(ctx context.Context, bioskop *entity.Bioskop) error
	GetAll(ctx context.Context) ([]entity.Bioskop, error)
	GetByID(ctx context.Context, id int) (*entity.Bioskop, error)
	Update(ctx context.Context, bioskop *entity.Bioskop) error
	Delete(ctx context.Context, id int) error
}

type bioskopRepository struct {
	db *sql.DB
}

func NewBioskopRepository(db *sql.DB) BioskopRepository {
	return &bioskopRepository{db: db}
}

func (r *bioskopRepository) Create(ctx context.Context, bioskop *entity.Bioskop) error {
	query := `INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRowContext(ctx, query, bioskop.Nama, bioskop.Lokasi, bioskop.Rating).Scan(&bioskop.ID)
}

func (r *bioskopRepository) GetAll(ctx context.Context) ([]entity.Bioskop, error) {
	query := `SELECT id, nama, lokasi, rating FROM bioskop ORDER BY id ASC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bioskops := []entity.Bioskop{}
	for rows.Next() {
		var b entity.Bioskop
		if err := rows.Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating); err != nil {
			return nil, err
		}
		bioskops = append(bioskops, b)
	}

	// Pengecekan error setelah iterasi selesai
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bioskops, nil
}

func (r *bioskopRepository) GetByID(ctx context.Context, id int) (*entity.Bioskop, error) {
	query := `SELECT id, nama, lokasi, rating FROM bioskop WHERE id = $1`
	var b entity.Bioskop
	err := r.db.QueryRowContext(ctx, query, id).Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("bioskop tidak ditemukan")
		}
		return nil, err
	}
	return &b, nil
}

func (r *bioskopRepository) Update(ctx context.Context, bioskop *entity.Bioskop) error {
	query := `UPDATE bioskop SET nama = $1, lokasi = $2, rating = $3 WHERE id = $4`
	result, err := r.db.ExecContext(ctx, query, bioskop.Nama, bioskop.Lokasi, bioskop.Rating, bioskop.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("bioskop tidak ditemukan")
	}
	return nil
}

func (r *bioskopRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM bioskop WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("bioskop tidak ditemukan")
	}
	return nil
}
