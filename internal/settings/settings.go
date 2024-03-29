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
	Debug     bool   `env:"DEBUG,default=false"`
	LogLevel  string `env:"LOG_LEVEL,default=info"`
	Port      string `env:"PORT,default=80"`
	StaticDir string `env:"STATIC_DIR"`
	DB        DB
	OAuth     OAuth
}

type OAuth struct {
	RedirectURLBase string `env:"OAUTH_REDIRECT_URL_BASE"`
	GitHubID        string `env:"GITHUB_CLIENT_ID"`
	GitHubSecret    string `env:"GITHUB_CLIENT_SECRET"`
	GoogleID        string `env:"GOOGLE_CLIENT_ID"`
	GoogleSecret    string `env:"GOOGLE_CLIENT_SECRET"`
}

func (o OAuth) RedirectURL(name string) string {
	return fmt.Sprintf("%s/api/oauth/callback/%s", o.RedirectURLBase, name)
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
