package appconfig

type AppConfig struct {
	Token  string
	SQLite struct {
		File string
	}
	GRPC struct {
		Port string
	}
}

type ConfigReader interface {
	GetAppConfig() *AppConfig
}
