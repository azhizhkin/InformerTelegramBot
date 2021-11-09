package infrastructure

import (
	"fmt"
	"log"

	"github.com/azhizhkin/informertelegrambot/internal/appconfig"
	"github.com/jinzhu/copier"

	"github.com/kkyr/fig"
)

type FigConfigReader struct {
	Token  string
	SQLite struct {
		File string
	}
	GRPC struct {
		Port string
	}
}

func (figConfig FigConfigReader) GetAppConfig() *appconfig.AppConfig {
	if err := fig.Load(&figConfig); err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	var appConfig appconfig.AppConfig
	err := copier.Copy(&appConfig, figConfig)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error creating config:%v", err))
	}
	return &appConfig
}
