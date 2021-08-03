package persistence

import (
	"fmt"

	"github.com/avenuesec/tilt-demo/pkg/config"
)

func BuildDsn(cfg *config.Configuration) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		cfg.PostgresUsername, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDatabase)
}
