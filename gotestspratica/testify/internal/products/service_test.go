package products_test

import (
	"errors"
	"testing"

	"github.com/batatinha123/bootcamp-meli-tests-pratica/internal/entities"
	"github.com/batatinha123/bootcamp-meli-tests-pratica/internal/products"
	mocks "github.com/batatinha123/bootcamp-meli-tests-pratica/tests/products"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllBySeller(t *testing.T) {
	var emptyProducts []entities.Product
	t.Run("deve retornar os sellers ao chamar a repository", func(t *testing.T) {
		mockProducts := []entities.Product{
			{
				ID:       "123",
				Name:     "Product 1",
				Price:    10.0,
				SellerId: "123",
			},
		}

		service, repoMock := createService(t)

		repoMock.Mock.On(
			"GetAllBySellerId",
			mock.AnythingOfType("string"),
		).Return(
			mockProducts,
			nil,
		)

		result, err := service.GetAllBySellerId("123")
		assert.Nil(t, err)
		assert.True(t, len(result) > 0)
	})

	t.Run("Deve retornar o erro caso ocorra um erro na chamada da repository", func(t *testing.T) {
		expectedErrorMessage := "erro qualquer"
		mockError := errors.New("erro qualquer")

		service, repoMock := createService(t)

		repoMock.Mock.On(
			"GetAllBySellerId",
			mock.AnythingOfType("string"),
		).Return(
			emptyProducts,
			mockError,
		)

		_, err := service.GetAllBySellerId("123")

		repoMock.Mock.AssertCalled(t, "GetAllBySellerId", "123")
		assert.NotNil(t, err)
		assert.Equal(t, expectedErrorMessage, err.Error())
	})
}

func createService(t *testing.T) (products.Service, *mocks.ProductsRepositoryMock) {
	t.Helper()
	repoMock := new(mocks.ProductsRepositoryMock)
	service := products.NewService(repoMock)
	return service, repoMock
}
