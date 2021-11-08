package main

import (
	"fmt"
	"log"
	"net"

	"github.com/AndreyZhizhkin/informertelegrambot/internal/grpcmessaging"
	"github.com/AndreyZhizhkin/informertelegrambot/internal/infrastructure"
	"github.com/AndreyZhizhkin/informertelegrambot/internal/telegrambotservice"
	"github.com/xlab/closer"
	"google.golang.org/grpc"
)

var sqliteRepo *infrastructure.SQLiteRepository
var gRPCservice *grpcmessaging.GRPCMessagingService

func main() {
	closer.Bind(cleanup)
	appConfig := infrastructure.FigConfigReader{}.GetAppConfig()
	sqliteRepo = infrastructure.NewSQLiteRepository(appConfig)
	tgBotService := telegrambotservice.NewTelegramBotService(appConfig, sqliteRepo)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", appConfig.GRPC.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gRPCservice = grpcmessaging.NewGRPCMessagingService(appConfig, tgBotService, grpc.NewServer(), &lis)
	go tgBotService.Run()
	go gRPCservice.Run()
	log.Println("InformerTelegramBot is running...")
	closer.Hold()
}

func cleanup() {
	sqliteRepo.DB.Close()
	gRPCservice.Stop()
	log.Println("InformerTelegramBot stopped")
}
