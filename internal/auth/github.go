package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"

	"github.com/jimmykodes/vehicle_maintenance/internal/settings"
)

type Github struct {
	config  *oauth2.Config
	baseURL string
}

func NewGithub(settings settings.OAuth) *Github {
	config := &oauth2.Config{
		ClientID:     settings.GitHubID,
		ClientSecret: settings.GitHubSecret,
		Endpoint:     endpoints.GitHub,
		RedirectURL:  settings.RedirectURL("github"),
		Scopes: []string{
			"read:user",
			"user:email",
		},
	}
	return &Github{
		config:  config,
		baseURL: "https://api.github.com",
	}
}
func (g Github) AuthCodeURL() (state string, url string, err error) {
	state, err = newState()
	if err != nil {
		return "", "", err
	}
	return state, g.config.AuthCodeURL(state), nil
}

func (g Github) exchange(code string) (*oauth2.Token, error) {
	return g.config.Exchange(context.Background(), code)
}

func (g Github) getUserInfo(token *oauth2.Token) (map[string]interface{}, error) {
	ctx := context.Background()
	c := g.config.Client(ctx, token)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/user", g.baseURL), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (g Github) GetUsername(code string) (username string, err error) {
	token, err := g.exchange(code)
	if err != nil {
		return "", err
	}
	userData, err := g.getUserInfo(token)
	if err != nil {
		return "", err
	}
	username, ok := userData["login"].(string)
	if !ok || username == "" {
		return "", ErrNoUsername
	}
	return username, nil
}

func newState() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
