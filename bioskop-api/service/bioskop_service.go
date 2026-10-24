package service

import (
	"bioskop-api/entity"
	"bioskop-api/repository"
	"context"
	"errors"
	"strings"
)

type BioskopService interface {
	CreateBioskop(ctx context.Context, input *entity.Bioskop) (*entity.Bioskop, error)
}

type bioskopService struct {
	repo repository.BioskopRepository
}

func NewBioskopService(repo repository.BioskopRepository) BioskopService {
	return &bioskopService{repo: repo}
}

func (s *bioskopService) CreateBioskop(ctx context.Context, input *entity.Bioskop) (*entity.Bioskop, error) {
	// Validasi bahwa Nama dan Lokasi tidak boleh kosong
	if strings.TrimSpace(input.Nama) == "" {
		return nil, errors.New("nama bioskop tidak boleh kosong")
	}
	if strings.TrimSpace(input.Lokasi) == "" {
		return nil, errors.New("lokasi bioskop tidak boleh kosong")
	}

	// Simpan ke repository
	err := s.repo.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	return input, nil
}
