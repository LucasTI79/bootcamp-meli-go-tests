package handler

import (
	"errors"
	"net/http"

	"github.com/batatinha123/bootcamp-meli-tests-pratica/internal/products"
	"github.com/batatinha123/bootcamp-meli-tests-pratica/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service products.Service
}

func NewProductHandler(service products.Service) *ProductHandler {
	return &ProductHandler{service}
}

func (h *ProductHandler) GetAllBySellerId(ctx *gin.Context) {
	sellerID := ctx.Query("seller_id")
	if sellerID == "" {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(
			http.StatusBadRequest,
			nil,
			"seller_id query param is required",
		))
		return
	}

	data, err := h.service.GetAllBySellerId(sellerID)

	if err != nil {
		switch {
		case errors.Is(err, products.ErrResourceNotFound):
			ctx.JSON(
				http.StatusNotFound,
				web.NewResponse(
					http.StatusNotFound,
					nil,
					err.Error(),
				),
			)
			return
		default:
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	if len(data) == 0 {
		ctx.Status(http.StatusNoContent)
		return
	}

	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, data, ""))
}
