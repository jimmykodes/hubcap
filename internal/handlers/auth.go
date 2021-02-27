package handlers

import (
	"database/sql"
	"errors"
	"github.com/jimmykodes/vehicle_maintenance/internal/auth"
	"github.com/jimmykodes/vehicle_maintenance/internal/dao"
	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
	"github.com/jimmykodes/vehicle_maintenance/internal/settings"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const errorRedirect = "/#error"

type Auth struct {
	logger  *zap.Logger
	userDAO dao.User
	oauth   *auth.Github
}

func NewAuth(logger *zap.Logger, userDAO dao.User, ghSettings settings.GitHubAuth) *Auth {
	return &Auth{logger: logger, userDAO: userDAO, oauth: auth.NewGithub(ghSettings)}
}

func (h Auth) Login(w http.ResponseWriter, r *http.Request) {
	state, url, err := h.oauth.AuthCodeURL()
	if err != nil {
		writeErrorResponse(w, h.logger, http.StatusInternalServerError, "")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "state",
		Value:    state,
		Expires:  time.Now().Add(time.Minute * 10),
		Secure:   false,
		HttpOnly: true,
	})
	http.Redirect(w, r, url, http.StatusFound)
}

// Callback is the destination from googleOAuth and handles creating the user and session
func (h Auth) Callback(w http.ResponseWriter, r *http.Request) {
	destination := h.oAuth2Handler(w, r)
	http.Redirect(w, r, destination, http.StatusFound)
}

// oAuth2Handler is the actual handler of the oauth2 return data
// returns redirect location as string
func (h Auth) oAuth2Handler(w http.ResponseWriter, r *http.Request) string {
	qp := map[string]string{
		"error": r.URL.Query().Get("error"),
		"code":  r.URL.Query().Get("code"),
		"state": r.URL.Query().Get("state"),
	}
	if qp["error"] != "" {
		if qp["error"] == "access_denied" {
			// user canceled just return to home screen
			return "/#canceled"
		}
		return errorRedirect
	}
	stateCookie, err := r.Cookie("state")
	if err != nil {
		h.logger.Error("error getting state", zap.Error(err))
		return errorRedirect
	}
	qpState := qp["state"]
	if qpState == "" {
		h.logger.Debug("no state from OAuth2")
		return errorRedirect
	}
	sessionState := stateCookie.Value
	if qpState != sessionState {
		h.logger.Debug("invalid state from OAuth2")
		return errorRedirect
	}

	if qp["code"] == "" {
		h.logger.Debug("No code supplied from OAuth2")
		return errorRedirect
	}

	token, err := h.oauth.Exchange(qp["code"])
	if err != nil {
		h.logger.Error("Error retrieving token", zap.Error(err))
		return errorRedirect
	}

	userData, err := h.oauth.GetUserInfo(token)
	if err != nil {
		h.logger.Error("Error getting user information", zap.Error(err))
		return errorRedirect
	}
	h.logger.Debug("user", zap.Any("data", userData))
	username, ok := userData["login"].(string)
	if !ok {
		h.logger.Error("Error getting user login")
		return errorRedirect
	}
	user, err := h.userDAO.GetFromUsername(r.Context(), username)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			h.logger.Error("error retrieving user from DB", zap.Error(err))
			return errorRedirect
		}
		if err := h.userDAO.Create(r.Context(), &dto.User{Username: username}); err != nil {
			h.logger.Error("error creating user", zap.Error(err), zap.String("username", username))
			return errorRedirect
		}
		user, err = h.userDAO.GetFromUsername(r.Context(), username)
		if err != nil {
			h.logger.Error("error getting user after user creation", zap.Error(err), zap.String("username", username))
			return errorRedirect
		}
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "apiKey",
		Value:    user.ApiKey,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		Secure:   false,
		HttpOnly: true,
	})
	return "/"
}
