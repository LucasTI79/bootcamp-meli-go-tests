package mocks

import (
	"github.com/batatinha123/bootcamp-meli-tests-pratica/internal/entities"
	"github.com/batatinha123/bootcamp-meli-tests-pratica/internal/products"
	"github.com/stretchr/testify/mock"
)

type ProductsServiceMock struct {
	mock.Mock
}

func NewProductsServiceMock() products.Service {
	return &ProductsServiceMock{}
}

func (p *ProductsServiceMock) GetAllBySellerId(sellerId string) ([]entities.Product, error) {
	args := p.Called(sellerId)
	// argumentos: [[]entities.Product{}, error]
	arg0, ok := args.Get(0).([]entities.Product)
	if !ok {
		return []entities.Product{}, args.Error(1)

	}
	return arg0, args.Error(1)
}

type ProductsRepositoryMock struct {
	mock.Mock
}

func (p *ProductsRepositoryMock) GetAllBySellerId(sellerId string) ([]entities.Product, error) {
	args := p.Called(sellerId)
	// argumentos: [[]entities.Product{}, error]
	return args.Get(0).([]entities.Product), args.Error(1)
}

func NewProductsRepositoryMock() products.Repository {
	return &ProductsRepositoryMock{}
}
