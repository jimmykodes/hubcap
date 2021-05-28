package handlers

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/jimmykodes/vehicle_maintenance/internal/auth"
	"github.com/jimmykodes/vehicle_maintenance/internal/dao"
	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
	"github.com/jimmykodes/vehicle_maintenance/internal/settings"
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
	user, err := h.getUser(r.Context(), username)
	if err != nil {
		h.logger.Error("error getting user", zap.Error(err))
		return errorRedirect
	}
	expires := time.Now().Add(time.Hour * 24)
	session, err := h.userDAO.CreateSession(r.Context(), user, expires)
	if err != nil {
		h.logger.Error("error creating user session", zap.Error(err), zap.String("username", username))
		return errorRedirect
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    session,
		Path:     "/",
		Expires:  expires,
		Secure:   false,
		HttpOnly: true,
	})
	return "/"
}

func (h Auth) getUser(ctx context.Context, username string) (*dto.User, error) {
	user, err := h.userDAO.GetFromUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		if err = h.userDAO.Create(ctx, &dto.User{Username: username}); err != nil {
			return nil, err
		}
		return h.getUser(ctx, username)
	}
	return user, nil
}
