package settings

import (
	"fmt"

	"github.com/Netflix/go-env"
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
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Database string `env:"DB_DATABASE"`
	URL      string `env:"DATABASE_URL"`
}

func (db DB) DSN() string {
	if db.URL != "" {
		return db.URL
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.Host, db.Port, db.User, db.Password, db.Database)
}
