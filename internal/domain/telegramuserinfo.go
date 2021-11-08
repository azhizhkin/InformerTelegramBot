package domain

type TelegramUserInfo struct {
	UserName string
	UserID   string
}

type TelegramUserInfoRepository interface {
	AddUserInfo(userInfo TelegramUserInfo) error
	GetUserID(userName string) (string, error)
}
