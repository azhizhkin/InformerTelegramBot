package infrastructure

import (
	"fmt"
	"log"

	"github.com/AndreyZhizhkin/Informer/core/internal/appcore/appservices"
	"github.com/jinzhu/copier"

	"github.com/kkyr/fig"
)

type FigConfigReader struct {
	MongoDB struct {
		URI                     string `fig:"URI" default:"mongodb://localhost:27017"`
		DB                      string `fig:"DB" default:"Informer"`
		MsgCollection           string `fig:"MsgCollection" default:"Messages"`
		SubscrCollection        string `fig:"SubscrCollection" default:"Subscribers"`
		TelegramUsersCollection string `fig:"TelegramUsersCollection" default:"TelegramUsers"`
	}
	HTTPAPIServer struct {
		Port string `fig:"Port" default:"8080"`
	}
	InfChannels struct {
		Mail struct {
			From string
			SMTP struct {
				From string
				Addr string
				Port string
				User string
				Pass string
			}
		}
		Telegram struct {
			Token string
		}
	}
	GRPC struct {
		Port string
	}
}

func (figConfig FigConfigReader) GetAppConfig() *appservices.AppConfig {
	if err := fig.Load(&figConfig); err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	var appConfig appservices.AppConfig
	err := copier.Copy(&appConfig, figConfig)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error creating config:%v", err))
	}
	return &appConfig
}
