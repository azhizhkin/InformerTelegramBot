package telegrambotservice

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/AndreyZhizhkin/informertelegrambot/internal/appconfig"
	"github.com/AndreyZhizhkin/informertelegrambot/internal/domain"

	tb "gopkg.in/tucnak/telebot.v2"
)

type TelegramBotService struct {
	repo domain.TelegramUserInfoRepository
	bot  *tb.Bot
}

func NewTelegramBotService(config *appconfig.AppConfig, repo domain.TelegramUserInfoRepository) *TelegramBotService {
	s := TelegramBotService{
		repo: repo,
	}
	var err error
	s.bot, err = tb.NewBot(tb.Settings{
		Token:  config.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}
	s.bot.Handle("/start", s.OnStartMessage)
	s.bot.Handle("/showuserid", s.ShowUserID)
	return &s
}

func (s TelegramBotService) Run() {
	s.bot.Start()
}

func (s TelegramBotService) OnStartMessage(m *tb.Message) {
	userInfo := domain.TelegramUserInfo{
		Name: m.Sender.Username,
		ID:   strconv.Itoa(m.Sender.ID),
	}
	err := s.repo.AddUserInfo(userInfo)
	if err != nil {
		s.bot.Send(m.Sender, fmt.Sprintf("Error adding user info:%v", err))
	} else {
		s.bot.Send(m.Sender, "Now you can receive messages from InformerBot!")
	}
}

func (s TelegramBotService) ShowUserID(m *tb.Message) {
	userID := s.repo.GetUserID(m.Sender.Username)
	if userID != "" {
		s.bot.Send(m.Sender, fmt.Sprintf("User ID is: %s", userID))
	} else {
		s.bot.Send(m.Sender, "User unknown")
	}
}