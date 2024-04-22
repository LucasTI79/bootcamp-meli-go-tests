package main

import (
	"fmt"

	"github.com/batatinha123/bootcamp-meli-tests/doubles"
	"github.com/batatinha123/bootcamp-meli-tests/service"
)

func main() {
	search := doubles.DummySearchEngine{}
	engine := service.NewEngine(search)

	fmt.Printf("A versão atual é %s", engine.GetVersion())
}
