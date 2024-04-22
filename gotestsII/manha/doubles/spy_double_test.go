package doubles_test

import (
	"testing"

	"github.com/batatinha123/bootcamp-meli-tests/doubles"
	"github.com/batatinha123/bootcamp-meli-tests/service"
	"github.com/stretchr/testify/assert"
)

func TestSearchByPhone(t *testing.T) {
	mySpySearchEngine := doubles.SpySearchEngine{SearchByPhoneWasCalled: false}
	engine := service.NewEngine(&mySpySearchEngine)
	phone := "12345678" // em nosso mecanismo, não realizaremos pesquisas se o telefone
	// tem um comprimento inferior a 5

	engine.SearchByPhone(phone)

	assert.True(t, mySpySearchEngine.SearchByPhoneWasCalled)
}

func TestSearchByPhoneNotCalled(t *testing.T) {
	mySpySearchEngine := doubles.SpySearchEngine{}
	engine := service.NewEngine(&mySpySearchEngine)

	engine.SearchByPhone("1234") // o telefone tem menos de 5 caracteres

	assert.False(t, mySpySearchEngine.SearchByPhoneWasCalled)
}
