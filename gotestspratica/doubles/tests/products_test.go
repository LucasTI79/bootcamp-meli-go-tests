package tests

import (
	"testing"

	"github.com/batatinha123/products-api/internal/products"
	"github.com/batatinha123/products-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDoubleGetAll(t *testing.T) {
	t.Run("Test_GetAll_OK", func(t *testing.T) {
		myStub := mocks.StubProducts{}
		myRepoStub := products.NewRepository(&myStub)
		myService := products.NewService(myRepoStub)

		result, err := myService.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
	})

	t.Run("Test_GetAll_No_Content", func(t *testing.T) {
		myStub := mocks.StubProductsNoContent{}
		myRepoStub := products.NewRepository(&myStub)
		myService := products.NewService(myRepoStub)

		result, err := myService.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, 0, len(result))
	})

	t.Run("Test_GetAll_Error", func(t *testing.T) {
		myStub := mocks.StubProductsError{}
		myRepoStub := products.NewRepository(&myStub)
		myService := products.NewService(myRepoStub)

		_, err := myService.GetAll()
		assert.NotNil(t, err)
		assert.Equal(t, "JSON unexpected caracter", err.Error())
	})
}
