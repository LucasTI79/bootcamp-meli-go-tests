package products

import (
	"github.com/batatinha123/bootcamp-meli-tests-pratica/internal/entities"
	"github.com/google/uuid"
)

type Repository interface {
	GetAllBySellerId(sellerId string) ([]entities.Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAllBySellerId(sellerId string) ([]entities.Product, error) {
	var productsList []entities.Product

	productsList = append(productsList, entities.Product{
		ID:       uuid.NewString(),
		Name:     "Product 1",
		Price:    10.0,
		SellerId: "70f5f19b-d473-4045-8820-a8e23f01b059",
	})

	return productsList, nil
}
