package main

import (
	"log"

	"github.com/AndreyZhizhkin/informertelegrambot/internal/infrastructure"
	"github.com/AndreyZhizhkin/informertelegrambot/internal/telegrambotservice"
	"github.com/xlab/closer"
)

var sqliteRepo *infrastructure.SQLiteRepository

func main() {
	closer.Bind(cleanup)
	appConfig := infrastructure.FigConfigReader{}.GetAppConfig()
	sqliteRepo = infrastructure.NewSQLiteRepository(appConfig)
	tgBotService := telegrambotservice.NewTelegramBotService(appConfig, sqliteRepo)
	go tgBotService.Run()
	log.Println("InformerTelegramBot is running...")
	closer.Hold()
}

func cleanup() {
	sqliteRepo.DB.Close()
	log.Println("InformerTelegramBot stopped")
}
