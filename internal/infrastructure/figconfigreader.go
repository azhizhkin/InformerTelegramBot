package infrastructure

import (
	"log"

	"github.com/azhizhkin/informertelegrambot/internal/appconfig"

	"github.com/kkyr/fig"
)

type FigConfigReader appconfig.AppConfig

func (figConfig FigConfigReader) GetAppConfig() *appconfig.AppConfig {
	if err := fig.Load(&figConfig); err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	return (*appconfig.AppConfig)(&figConfig)
}
