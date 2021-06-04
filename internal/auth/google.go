package auth

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
	oauth2api "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"

	"github.com/jimmykodes/vehicle_maintenance/internal/settings"
)

type Google struct {
	config *oauth2.Config
}

func NewGoogle(settings settings.OAuth) *Google {
	return &Google{config: &oauth2.Config{
		ClientID:     settings.GoogleID,
		ClientSecret: settings.GoogleSecret,
		Endpoint:     endpoints.Google,
		RedirectURL:  settings.RedirectURL("google"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // user info scope
		},
	}}
}
func (g Google) AuthCodeURL() (string, string, error) {
	state, err := newState()
	if err != nil {
		return "", "", err
	}
	return state, g.config.AuthCodeURL(state), nil
}

func (g Google) exchange(code string) (*oauth2.Token, error) {
	return g.config.Exchange(context.Background(), code)
}

func (g Google) getUserInfo(token *oauth2.Token) (*oauth2api.Userinfo, error) {
	ctx := context.Background()
	userService, err := oauth2api.NewService(ctx, option.WithTokenSource(g.config.TokenSource(ctx, token)))
	if err != nil {
		return nil, err
	}
	userInfoService := oauth2api.NewUserinfoV2MeService(userService)
	return userInfoService.Get().Do()
}

func (g Google) GetUsername(code string) (username string, err error) {
	token, err := g.exchange(code)
	if err != nil {
		return "", err
	}
	userData, err := g.getUserInfo(token)
	if err != nil {
		return "", err
	}
	return userData.Email, nil
}
