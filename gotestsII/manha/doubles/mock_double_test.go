package doubles_test

import (
	"testing"

	"github.com/batatinha123/bootcamp-meli-tests/doubles"
	"github.com/batatinha123/bootcamp-meli-tests/service"
	"github.com/stretchr/testify/assert"
)

func TestSearchByNameMocked(t *testing.T) {
	myMockSearchEngine := doubles.MockSearchEngine{}
	engine := service.NewEngine(&myMockSearchEngine)
	expectedPhone := "12345678"

	result := engine.SearchByName("Nacho")

	assert.Equal(t, expectedPhone, result)
	assert.True(t, myMockSearchEngine.SearchByNameWasCalled)
}

func TestSearchByNameMockedNotCalled(t *testing.T) {
	myMockSearchEngine := doubles.MockSearchEngine{}
	engine := service.NewEngine(&myMockSearchEngine)
	expectedPhone := ""

	result := engine.SearchByName("Nac")

	assert.Equal(t, expectedPhone, result)
	assert.False(t, myMockSearchEngine.SearchByNameWasCalled)
}
