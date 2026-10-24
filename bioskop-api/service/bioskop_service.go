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
	GetAllBioskop(ctx context.Context) ([]entity.Bioskop, error)
	GetBioskopByID(ctx context.Context, id int) (*entity.Bioskop, error)
	UpdateBioskop(ctx context.Context, id int, input *entity.Bioskop) (*entity.Bioskop, error)
	DeleteBioskop(ctx context.Context, id int) error
}

type bioskopService struct {
	repo repository.BioskopRepository
}

func NewBioskopService(repo repository.BioskopRepository) BioskopService {
	return &bioskopService{repo: repo}
}

func (s *bioskopService) CreateBioskop(ctx context.Context, input *entity.Bioskop) (*entity.Bioskop, error) {
	if strings.TrimSpace(input.Nama) == "" {
		return nil, errors.New("nama bioskop tidak boleh kosong")
	}
	if strings.TrimSpace(input.Lokasi) == "" {
		return nil, errors.New("lokasi bioskop tidak boleh kosong")
	}

	err := s.repo.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (s *bioskopService) GetAllBioskop(ctx context.Context) ([]entity.Bioskop, error) {
	return s.repo.GetAll(ctx)
}

func (s *bioskopService) GetBioskopByID(ctx context.Context, id int) (*entity.Bioskop, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *bioskopService) UpdateBioskop(ctx context.Context, id int, input *entity.Bioskop) (*entity.Bioskop, error) {
	if strings.TrimSpace(input.Nama) == "" {
		return nil, errors.New("nama bioskop tidak boleh kosong")
	}
	if strings.TrimSpace(input.Lokasi) == "" {
		return nil, errors.New("lokasi bioskop tidak boleh kosong")
	}

	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	input.ID = id
	err = s.repo.Update(ctx, input)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (s *bioskopService) DeleteBioskop(ctx context.Context, id int) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}
