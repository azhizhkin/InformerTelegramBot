package main

import (
	"fmt"
	"log"
	"net"

	"github.com/AndreyZhizhkin/informertelegrambot/internal/appconfig"
	"github.com/AndreyZhizhkin/informertelegrambot/internal/domain"
	"github.com/AndreyZhizhkin/informertelegrambot/internal/grpcmessaging"
	"github.com/AndreyZhizhkin/informertelegrambot/internal/infrastructure"
	"github.com/AndreyZhizhkin/informertelegrambot/internal/telegrambotservice"

	_ "github.com/mattn/go-sqlite3"
	"github.com/xlab/closer"
	"go.uber.org/dig"
	"google.golang.org/grpc"
)

var (
	container  *dig.Container
	sqliteRepo *infrastructure.SQLiteRepository
	gRPCMsg    *grpcmessaging.GRPCMessagingService
)

func main() {
	closer.Bind(cleanup)
	buildContainer()
	err := container.Invoke(func(tgUserInfoRepo domain.TelegramUserInfoRepository) {
		sqliteRepo = tgUserInfoRepo.(*infrastructure.SQLiteRepository)
	})
	if err != nil {
		panic(err)
	}
	err = container.Invoke(func(tgBotService *telegrambotservice.TelegramBotService) {
		go tgBotService.Run()
	})
	if err != nil {
		panic(err)
	}
	err = container.Invoke(func(grpcMsg *grpcmessaging.GRPCMessagingService) {
		go grpcMsg.Run()
	})
	if err != nil {
		panic(err)
	}
	log.Println("InformerTelegramBot is running...")
	closer.Hold()
}

func buildContainer() {
	container = dig.New()
	container.Provide(infrastructure.FigConfigReader{}.GetAppConfig)
	container.Provide(infrastructure.NewSQLiteRepository, dig.As(new(domain.TelegramUserInfoRepository)))
	container.Provide(telegrambotservice.NewTelegramBotService)
	container.Provide(func(appConfig *appconfig.AppConfig) *net.Listener {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", appConfig.GRPC.Port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		return &lis
	})
	container.Provide(grpcmessaging.NewGRPCMessagingService)
	container.Provide(grpc.NewServer)
}

func cleanup() {
	sqliteRepo.DB.Close()
	gRPCMsg.Stop()
	log.Println("InformerTelegramBot stopped")
}
