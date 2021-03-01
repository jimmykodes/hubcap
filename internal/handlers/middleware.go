package handlers

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
	"net/http"
	"time"

	"github.com/jimmykodes/vehicle_maintenance/internal/dao"
	"go.uber.org/zap"
)

type MiddlewareFunc func(h http.HandlerFunc) http.HandlerFunc

func NewMiddleware(logger *zap.Logger, userDAO dao.User) *Middleware {
	m := &Middleware{logger: logger, userDAO: userDAO}
	m.Standard = []MiddlewareFunc{
		m.Log,
		m.CORS,
		m.Auth,
	}
	return m
}

type Middleware struct {
	logger   *zap.Logger
	userDAO  dao.User
	Standard []MiddlewareFunc
}

func (m Middleware) writeError(w http.ResponseWriter, status int, msg string) {
	writeErrorResponse(w, m.logger, status, msg)
}

func (m *Middleware) Reduce(h http.HandlerFunc, mf ...MiddlewareFunc) http.HandlerFunc {
	if len(mf) < 1 {
		return h
	}
	wrapped := h
	// loop in reverse to preserve middleware order
	for i := len(mf) - 1; i >= 0; i-- {
		wrapped = mf[i](wrapped)
	}
	return wrapped
}

func (m *Middleware) CORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Accept, Accept-Encoding, Accept-Charset, Content-Type, Content-Length",
		)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		h.ServeHTTP(w, r)
	}
}
func (m *Middleware) Log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.logger.Info("received request", zap.String("path", r.URL.String()))
		h.ServeHTTP(w, r)
		m.logger.Info("completed request", zap.String("path", r.URL.String()))
	}
}

func (m *Middleware) Auth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			user *dto.User
			err  error
		)
		apiKey := r.Header.Get("X-API-KEY")
		sessionKey := getSessionKey(r)
		switch {
		case apiKey != "":
			user, err = m.userDAO.GetFromApiKey(r.Context(), apiKey)
		case sessionKey != "":
			user, err = m.userDAO.GetFromSession(r.Context(), sessionKey, time.Now().Unix())
		default:
			m.writeError(w, http.StatusUnauthorized, "missing authentication")
			return
		}
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				m.writeError(w, http.StatusUnauthorized, "unauthorized")
				return
			}
			m.writeError(w, http.StatusInternalServerError, "")
			return
		}
		ctx := context.WithValue(r.Context(), userIDKey, user.ID)
		ctx = context.WithValue(ctx, userKey, user)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
}

func getSessionKey(r *http.Request) string {
	cookie, err := r.Cookie("session")
	if err != nil {
		// only error possible is ErrNoCookie, so just return an empty session key
		return ""
	}
	return cookie.Value
}
