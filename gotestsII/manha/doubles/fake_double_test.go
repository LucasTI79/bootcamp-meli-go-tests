package doubles_test

import (
	"testing"

	"github.com/batatinha123/bootcamp-meli-tests/doubles"
	"github.com/batatinha123/bootcamp-meli-tests/service"
	"github.com/stretchr/testify/assert"
)

func TestSearchByNameFaked(t *testing.T) {
	myFakeSearchEngine := doubles.FakeSearchEngine{}
	engine := service.NewEngine(&myFakeSearchEngine)

	// table test
	testValues := map[string]string{"Nacho": "123456", "Nico": "234567"}

	for key := range testValues {
		engine.AddEntry(key, testValues[key])
	}

	resultNacho := engine.SearchByName("Nacho")
	resultNico := engine.SearchByName("Nico")
	resultPhone := engine.SearchByPhone("123456")

	assert.Equal(t, testValues["Nacho"], resultNacho)
	assert.Equal(t, testValues["Nico"], resultNico)
	assert.Equal(t, "Nacho", resultPhone)

}
