package doubles_test

import (
	"testing"

	"github.com/batatinha123/bootcamp-meli-tests/doubles"
	"github.com/batatinha123/bootcamp-meli-tests/service"
	"github.com/stretchr/testify/assert"
)

func TestFindByName(t *testing.T) {
	myStubSearchEngine := doubles.StubSearchEngine{}
	engine := service.NewEngine(myStubSearchEngine)
	expectedPhone := "12345678"

	result := engine.SearchByName("Pepe")

	assert.Equal(t, expectedPhone, result)
}
