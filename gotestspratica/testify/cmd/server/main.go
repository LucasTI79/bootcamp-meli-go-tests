package main

import (
	"github.com/batatinha123/bootcamp-meli-tests-pratica/cmd/server/router"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	router.MapRoutes(server)

	server.Run()
}
