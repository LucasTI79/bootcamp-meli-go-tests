package mocks

import (
	"encoding/json"

	"github.com/batatinha123/products-api/internal/entities"
)

type StubProductsNoContent struct {
}

func (s *StubProductsNoContent) Write(data interface{}) error {
	return nil
}

func (s *StubProductsNoContent) Read(data interface{}) error {
	var products []entities.Product
	jsonData, _ := json.Marshal(products)
	return json.Unmarshal(jsonData, &data)
}
