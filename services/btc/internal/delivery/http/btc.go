package http

import (
	"net/http"

	"github.com/avenuesec/tilt-demo/services/btc/internal/service"
	"github.com/gin-gonic/gin"
)

type BTCController struct {
	service *service.BTCPriceService
}

func NewBTCController(svc *service.BTCPriceService) *BTCController {
	controller := &BTCController{
		service: svc,
	}

	return controller
}

func (c *BTCController) RegisterRouter(engine *gin.Engine) {
	engine.GET("/btc", func(ctx *gin.Context) {
		res, err := c.service.Price(ctx)

		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, res)
	})
}
