package product

import (
	"net/http"
	"rest-api-fga/service/product"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *product.Service
}

func NewHandler(svc *product.Service) *Handler {
	return &Handler{
		service: svc,
	}
}

func (h *Handler) GetAllProducts(ctx *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}
