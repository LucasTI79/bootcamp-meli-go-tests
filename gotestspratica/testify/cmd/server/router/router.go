package router

import "github.com/gin-gonic/gin"

func MapRoutes(r *gin.Engine) {
	rg := r.Group("/api/v1")
	buildProductsRoutes(rg)
}

func buildProductsRoutes(rg *gin.RouterGroup) {
	productsGroup := rg.Group("/products")
	productsGroup.GET("/", nil)
}
