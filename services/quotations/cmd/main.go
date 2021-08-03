package main

import (
	"log"

	"github.com/avenuesec/tilt-demo/pkg/config"
	"github.com/avenuesec/tilt-demo/pkg/persistence"
	"github.com/avenuesec/tilt-demo/services/quotations/internal/delivery/http"
	"github.com/avenuesec/tilt-demo/services/quotations/internal/repository"
	"github.com/avenuesec/tilt-demo/services/quotations/internal/service"
	"github.com/caarlos0/env/v6"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	cfg = &config.Configuration{}
)

func init() {
	godotenv.Load("local.env")
	env.Parse(cfg)
}

func main() {
	server := gin.Default()

	dsn := persistence.BuildDsn(cfg)

	log.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	history, err := repository.NewHistoryRepository(db)

	if err != nil {
		panic(err)
	}

	factory := service.NewQuotationFactory(cfg)
	svc := service.NewQuotationService(factory, history)
	controller := http.NewQuotationController(svc)

	controller.RegisterRouter(server)

	server.Run(cfg.QuotationsHost)
}
