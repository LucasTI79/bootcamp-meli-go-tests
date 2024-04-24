package products

import (
	"database/sql"
	"errors"

	"github.com/batatinha123/bootcamp-meli-tests-pratica/internal/entities"
)

type Service interface {
	GetAllBySellerId(sellerId string) ([]entities.Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAllBySellerId(sellerId string) ([]entities.Product, error) {

	products, err := s.repository.GetAllBySellerId(sellerId)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrResourceNotFound
		}
		return nil, err
	}

	return products, nil
}
