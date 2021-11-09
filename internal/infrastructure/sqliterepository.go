package infrastructure

import (
	"database/sql"

	"github.com/AndreyZhizhkin/informertelegrambot/internal/appconfig"
	"github.com/AndreyZhizhkin/informertelegrambot/internal/domain"
)

type SQLiteRepository struct {
	DB *sql.DB
}

func NewSQLiteRepository(config *appconfig.AppConfig) *SQLiteRepository {
	var err error
	rep := SQLiteRepository{}
	rep.DB, err = sql.Open("sqlite3", config.SQLite.File)
	if err != nil {
		panic(err)
	}
	schemaSQL := `
 	CREATE TABLE IF NOT EXISTS tgusersinfo (
	name text UNIQUE,
	ID text
	);
	CREATE INDEX IF NOT EXISTS tgusersinfo_name ON tgusersinfo(name);`
	if _, err = rep.DB.Exec(schemaSQL); err != nil {
		panic(err)
	}

	return &rep
}

func (rep SQLiteRepository) AddUserInfo(userInfo domain.TelegramUserInfo) error {
	_, err := rep.DB.Exec("insert into tgusersinfo (name, ID) values ($1, $2)", userInfo.Name, userInfo.ID)
	if err != nil {
		return err
	}
	return nil
}

func (rep SQLiteRepository) GetUserID(userName string) (userID string) {
	row := rep.DB.QueryRow("select ID from tgusersinfo where name = $1", userName)
	err := row.Scan(&userID)
	if err != nil {
		panic(err)
	}
	return userID
}
