package http

import (
	"net/http"

	"github.com/avenuesec/tilt-demo/services/quotations/internal/service"
	"github.com/gin-gonic/gin"
)

type QuotationController struct {
	service *service.QuotationService
}

func NewQuotationController(svc *service.QuotationService) *QuotationController {
	controller := &QuotationController{
		service: svc,
	}

	return controller
}

func (c *QuotationController) RegisterRouter(engine *gin.Engine) {
	engine.GET("/quotations/:currency", func(ctx *gin.Context) {
		currency := service.Currency(ctx.Param("currency"))

		res, err := c.service.CurrencyPrice(ctx, currency)

		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, res)
	})
}
