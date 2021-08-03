package config

type Configuration struct {
	QuotationsHost   string `env:"QUOTATIONS_HOST"`
	BtcHost          string `env:"BTC_HOST"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     int    `env:"POSTGRES_PORT"`
	PostgresDatabase string `env:"POSTGRES_DATABASE"`
	PostgresUsername string `env:"POSTGRES_USERNAME"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PtaxURL          string `env:"PTAX_URL"`
	CoindekURL       string `env:"COINDECK_URL"`
}
