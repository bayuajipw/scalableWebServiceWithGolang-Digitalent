package handler

import (
	"rest-api-fga/handler/product"

	"github.com/gin-gonic/gin"
)

func NewHttpServer(productHandler *product.Handler) {
	r := gin.Default()

	r.GET("/products", productHandler.GetAllProducts)

	r.Run(":8000")
}
