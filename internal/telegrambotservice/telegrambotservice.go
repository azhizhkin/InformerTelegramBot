package telegrambotservice

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/azhizhkin/informertelegrambot/internal/appconfig"
	"github.com/azhizhkin/informertelegrambot/internal/domain"

	tb "gopkg.in/tucnak/telebot.v2"
)

type TelegramBotService struct {
	repo domain.TelegramUserInfoRepository
	bot  *tb.Bot
}

type telegramRecipient string

func (tu telegramRecipient) Recipient() string {
	return string(tu)
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
	if s.repo.GetUserID(userInfo.Name) != "" {
		s.bot.Send(m.Sender, "You can receive messages from InformerBot!")
		return
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

func (s TelegramBotService) SendMessage(userName string, text string) error {
	userID := s.repo.GetUserID(userName)
	_, err := s.bot.Send(telegramRecipient(userID), text)
	return err
}
