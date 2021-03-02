package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/jimmykodes/vehicle_maintenance/internal/dao"
	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
)

type User struct {
	logger  *zap.Logger
	userDAO dao.User
}

func NewUser(logger *zap.Logger, userDAO dao.User) *User {
	localLogger := logger.With(zap.String("handler", "user"))
	return &User{logger: localLogger, userDAO: userDAO}
}

func (h User) Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(userKey).(*dto.User)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		h.logger.Error("error writing user", zap.Error(err), zap.String("username", user.Username))
		w.WriteHeader(http.StatusInternalServerError)
	}
}
