package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/batatinha123/bootcamp-meli-tests-pratica/cmd/server/handler"
	"github.com/batatinha123/bootcamp-meli-tests-pratica/internal/entities"
	"github.com/batatinha123/bootcamp-meli-tests-pratica/internal/products"
	"github.com/batatinha123/bootcamp-meli-tests-pratica/pkg/testutil"
	"github.com/batatinha123/bootcamp-meli-tests-pratica/pkg/web"
	mocks "github.com/batatinha123/bootcamp-meli-tests-pratica/tests/products"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	t.Run("deve retornar status code 400 se o query param seller_id não for passado", func(t *testing.T) {
		var resp web.Response
		server, _ := InitServerWithGetProductsRoute(t)
		request, response := testutil.MakeRequest(http.MethodGet, GetAllProductsBySellerId, "")

		server.ServeHTTP(response, request)
		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "seller_id query param is required", resp.Error)
	})
	t.Run("deve retornar status code 404 se o seller não existir", func(t *testing.T) {
		var resp web.Response
		server, service := InitServerWithGetProductsRoute(t)
		service.On("GetAllBySellerId", "123").Return([]entities.Product{}, products.ErrResourceNotFound)

		request, response := testutil.MakeRequest(http.MethodGet, GetAllProductsBySellerId, "")
		DefineQueryParamSellerId(t, request)

		server.ServeHTTP(response, request)

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("deve retornar status code 204 se o seller não tiver produtos", func(t *testing.T) {
		server, service := InitServerWithGetProductsRoute(t)
		service.On("GetAllBySellerId", "123").Return([]entities.Product{}, nil)

		request, response := testutil.MakeRequest(http.MethodGet, GetAllProductsBySellerId, "")
		DefineQueryParamSellerId(t, request)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})
	t.Run("deve retornar status code 500, caso o método GetAllBySeller da repository retorne um erro", func(t *testing.T) {
		server, service := InitServerWithGetProductsRoute(t)
		service.On("GetAllBySellerId", "123").Return([]entities.Product{}, errors.New("erro desconhecido"))

		request, response := testutil.MakeRequest(http.MethodGet, GetAllProductsBySellerId, "")
		DefineQueryParamSellerId(t, request)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
	t.Run("deve retornar status code 200, com os produtos como resposta", func(t *testing.T) {
		t.Skip()
	})
}

const (
	GetAllProductsBySellerId = "/api/v1/products/"
)

func InitServerWithGetProductsRoute(t *testing.T) (*gin.Engine, *mocks.ProductsServiceMock) {
	t.Helper()
	server := testutil.CreateServer()
	mockService := new(mocks.ProductsServiceMock)
	handler := handler.NewProductHandler(mockService)
	server.GET(GetAllProductsBySellerId, handler.GetAllBySellerId)
	return server, mockService
}

func DefineQueryParamSellerId(t *testing.T, request *http.Request) {

	t.Helper()
	// Estamos pegando a query param da url
	// new URLSearchParams
	query := request.URL.Query() // /api/v1/products
	query.Add("seller_id", "123")
	// A função encode deixa os query params que foram definidos, no formato que o navegador entenda
	request.URL.RawQuery = query.Encode()
	// /api/v1/products?seller_id=123
}
