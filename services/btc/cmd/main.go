package main

import (
	"net/http"

	"github.com/avenuesec/tilt-demo/pkg/config"
	delivery "github.com/avenuesec/tilt-demo/services/btc/internal/delivery/http"
	"github.com/avenuesec/tilt-demo/services/btc/internal/repository"
	"github.com/avenuesec/tilt-demo/services/btc/internal/service"
	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
)

var (
	cfg = &config.Configuration{}
)

func init() {
	env.Parse(cfg)
}

func main() {
	server := gin.Default()

	coindeckRepo := repository.NewCointdeckRepository(cfg.CoindekURL, http.DefaultClient)
	quotationRepo := repository.NewQuotationRepository(cfg.QuotationsHost, http.DefaultClient)
	svc := service.NewBTCPriceService(quotationRepo, coindeckRepo)
	controller := delivery.NewBTCController(svc)

	controller.RegisterRouter(server)

	server.Run(cfg.BtcHost)
}
