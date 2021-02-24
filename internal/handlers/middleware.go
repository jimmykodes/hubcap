package handlers

import (
	"context"
	"net/http"

	"github.com/jimmykodes/vehicle_maintenance/internal/dao"
	"go.uber.org/zap"
)

type MiddlewareFunc func(h http.Handler) http.Handler

func NewMiddleware(logger *zap.Logger, userDAO dao.User) *Middleware {
	return &Middleware{logger: logger, userDAO: userDAO}
}

type Middleware struct {
	logger  *zap.Logger
	userDAO dao.User
}

func (m Middleware) writeError(w http.ResponseWriter, status int, msg string) {
	writeErrorResponse(w, m.logger, status, msg)
}

func (m *Middleware) Reduce(h http.Handler, mf ...MiddlewareFunc) http.Handler {
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

func (m *Middleware) CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})
}

func (m *Middleware) Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey == "" {
			m.writeError(w, http.StatusUnauthorized, "missing api key")
			return
		}
		user, err := m.userDAO.GetFromApiKey(r.Context(), apiKey)
		if err != nil {
			m.writeError(w, http.StatusInternalServerError, "")
			return
		}
		ctx := context.WithValue(r.Context(), userIDKey, user.ID)
		ctx = context.WithValue(ctx, userKey, user)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
