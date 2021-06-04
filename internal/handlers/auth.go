package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"

	"github.com/jimmykodes/vehicle_maintenance/internal/auth"
	"github.com/jimmykodes/vehicle_maintenance/internal/dao"
	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
	"github.com/jimmykodes/vehicle_maintenance/internal/settings"
)

const (
	canceledRedirect  = "/#canceled"
	errorRedirect     = "/#error"
	stateCookieName   = "state"
	SessionCookieName = "session"
)

type Auth struct {
	logger      *zap.Logger
	userDAO     dao.User
	githubOAuth *auth.Github
	googleOAuth *auth.Google
}

func NewAuth(logger *zap.Logger, userDAO dao.User, oauthSettings settings.OAuth) *Auth {
	return &Auth{
		logger:      logger,
		userDAO:     userDAO,
		githubOAuth: auth.NewGithub(oauthSettings),
		googleOAuth: auth.NewGoogle(oauthSettings),
	}
}

func (h Auth) Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := vars["service"]
	var oauth auth.OAuth
	switch vars["service"] {
	case auth.GoogleService:
		oauth = h.googleOAuth
	case auth.GitHubService:
		oauth = h.githubOAuth
	default:
		h.logger.Error("invalid login service", zap.String("service", service))
		writeErrorResponse(w, h.logger, http.StatusBadRequest, "invalid oauth service")
		return
	}

	state, url, err := oauth.AuthCodeURL()
	if err != nil {
		writeErrorResponse(w, h.logger, http.StatusInternalServerError, "")
		return
	}
	h.logger.Debug("setting state cookie", zap.String("cookie name", stateCookieName), zap.String("state", state))
	http.SetCookie(w, &http.Cookie{
		Name:     stateCookieName,
		Value:    state,
		Expires:  time.Now().Add(time.Minute * 10),
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	})
	http.Redirect(w, r, url, http.StatusFound)
}

// Callback is the destination from oauth handlers and handles creating the user and session
func (h Auth) Callback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := vars["service"]
	if service == "" {
		h.logger.Error("missing service")
		http.Redirect(w, r, errorRedirect, http.StatusFound)
		return
	}
	localLogger := h.logger.With(zap.String("service", service))
	stateCookie, err := r.Cookie(stateCookieName)
	if err != nil {
		localLogger.Error("error getting state cookie", zap.Error(err))
		http.Redirect(w, r, errorRedirect, http.StatusFound)
		return
	}
	resp := auth.NewResponse(r.URL.Query())
	if err := resp.Validate(stateCookie.Value); err != nil {
		if errors.Is(err, auth.ErrCanceled) {
			http.Redirect(w, r, canceledRedirect, http.StatusFound)
			return
		}
		localLogger.Error("invalid oauth response", zap.Error(err))
		http.Redirect(w, r, errorRedirect, http.StatusFound)
		return
	}
	ctx := r.Context()
	var oauth auth.OAuth
	switch service {
	case auth.GoogleService:
		oauth = h.googleOAuth
	case auth.GitHubService:
		oauth = h.githubOAuth
	default:
		http.Redirect(w, r, errorRedirect, http.StatusFound)
		return
	}
	username, err := oauth.GetUsername(resp.Code)
	if err != nil {
		http.Redirect(w, r, errorRedirect, http.StatusFound)
		return
	}
	user, err := h.getOrCreateUser(ctx, username)
	if err != nil {
		http.Redirect(w, r, errorRedirect, http.StatusFound)
		return
	}
	expires := time.Now().Add(time.Hour * 24)
	session, err := h.userDAO.CreateSession(ctx, user, expires)
	if err != nil {
		localLogger.Error("error creating session", zap.Error(err))
		http.Redirect(w, r, errorRedirect, http.StatusFound)
		return
	}
	h.setSessionCookie(w, session, expires)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h Auth) LogOut(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie(SessionCookieName)
	if err != nil {
		h.logger.Debug("error getting session cookie, nothing to do")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if err := h.userDAO.DeleteSession(r.Context(), sessionCookie.Value); err != nil {
		h.logger.Error("error deleting session", zap.String("session", sessionCookie.Value), zap.Error(err))
	}
	h.setSessionCookie(w, "", time.Now())
	http.Redirect(w, r, "/", http.StatusFound)
}

// getOrCreateUser will lookup the user by username, if one does not exist, it will be created
func (h Auth) getOrCreateUser(ctx context.Context, username string) (*dto.User, error) {
	user, err := h.userDAO.GetFromUsername(ctx, username)
	if errors.Is(err, pgx.ErrNoRows) {
		if err = h.userDAO.Create(ctx, &dto.User{Username: username}); err != nil {
			return nil, err
		}
		return h.getOrCreateUser(ctx, username)
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (h Auth) setSessionCookie(w http.ResponseWriter, session string, expires time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    session,
		Path:     "/",
		Expires:  expires,
		Secure:   false,
		HttpOnly: true,
	})
}
