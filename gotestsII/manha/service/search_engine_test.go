package service_test

import (
	"testing"

	"github.com/batatinha123/bootcamp-meli-tests/doubles"
	"github.com/batatinha123/bootcamp-meli-tests/service"
	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	myDummyDB := doubles.DummySearchEngine{}
	engine := service.NewEngine(myDummyDB)
	expectedVersion := "1.0"

	result := engine.GetVersion()

	assert.Equal(t, expectedVersion, result)
}
