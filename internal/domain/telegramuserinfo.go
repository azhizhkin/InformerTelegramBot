package domain

type TelegramUserInfo struct {
	Name string
	ID   string
}

type TelegramUserInfoRepository interface {
	AddUserInfo(userInfo TelegramUserInfo) error
	GetUserID(userName string) string
}
