package settings

import (
	"fmt"

	"github.com/Netflix/go-env"
	"github.com/go-sql-driver/mysql"
)

func NewSettings() (*Settings, error) {
	settings := &Settings{}
	_, err := env.UnmarshalFromEnviron(settings)
	if err != nil {
		return nil, err
	}
	return settings, nil
}

type Settings struct {
	Debug      bool   `env:"DEBUG,default=false"`
	LogLevel   string `env:"LOG_LEVEL,default=info"`
	Port       string `env:"PORT,default=80"`
	StaticDir  string `env:"STATIC_DIR"`
	DB         DB
	GitHubAuth GitHubAuth
}

type GitHubAuth struct {
	ID          string `env:"GITHUB_CLIENT_ID"`
	Secret      string `env:"GITHUB_CLIENT_SECRET"`
	RedirectURL string `env:"GITHUB_REDIRECT_URL"`
}

type DB struct {
	DriveName string `env:"DB_DRIVER_NAME,default=mysql"`
	Host      string `env:"DB_HOST"`
	Port      int    `env:"DB_PORT"`
	User      string `env:"DB_USER"`
	Password  string `env:"DB_PASSWORD"`
	Database  string `env:"DB_DATABASE"`
	Dns       string `env:"DB_DNS"`
}

func (db DB) addr() string {
	if db.Port != 0 {
		return fmt.Sprintf("%s:%d", db.Host, db.Port)
	}
	return db.Host
}

func (db DB) DNS() string {
	if db.Dns != "" {
		return db.Dns
	}
	switch db.DriveName {
	case "mysql":
		return (&mysql.Config{
			User:                 db.User,
			Passwd:               db.Password,
			Net:                  "tcp",
			Addr:                 db.addr(),
			AllowNativePasswords: true,
		}).FormatDSN()
	case "pg":
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.Host, db.Port, db.User, db.Password, db.Database)
	default:
		return ""
	}
}
