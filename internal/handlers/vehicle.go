package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jimmykodes/vehicle_maintenance/internal/dao"
	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
	"go.uber.org/zap"
)

type Vehicle struct {
	logger         *zap.Logger
	serviceTypeDAO dao.Vehicle
}

func NewVehicle(logger *zap.Logger, serviceDAO dao.Service) *Service {
	localLogger := logger.With(zap.String("handler", "vehicle"))
	return &Service{logger: localLogger, serviceDAO: serviceDAO}
}

func (h Vehicle) Detail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		h.logger.Error("error parsing ID", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userID := r.Context().Value(userIDKey).(int64)
	switch r.Method {
	case http.MethodGet:
		h.get(w, r, id, userID)
	case http.MethodPut:
		h.update(w, r, id, userID)
	case http.MethodDelete:
		h.delete(w, r, id, userID)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Vehicle) List(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(userIDKey).(int64)
	switch r.Method {
	case http.MethodGet:
		h.list(w, r, userID)
	case http.MethodPost:
		h.create(w, r, userID)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Vehicle) create(w http.ResponseWriter, r *http.Request, userID int64) {
	service := &dto.Vehicle{}
	if err := json.NewDecoder(r.Body).Decode(service); err != nil {
		h.logger.Error("error decoding json", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	service.UserID = userID
	if err := h.serviceTypeDAO.Create(r.Context(), service); err != nil {
		h.logger.Error("error calling create", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (h Vehicle) list(w http.ResponseWriter, r *http.Request, userID int64) {

}
func (h Vehicle) get(w http.ResponseWriter, r *http.Request, id, userID int64) {
	service, err := h.serviceTypeDAO.Get(r.Context(), id, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		h.logger.Error("error calling get", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(service); err != nil {
		h.logger.Error("error writing data", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func (h Vehicle) update(w http.ResponseWriter, r *http.Request, id, userID int64) {
	var service *dto.Vehicle
	if err := json.NewDecoder(r.Body).Decode(service); err != nil {
		h.logger.Error("error decoding json", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := h.serviceTypeDAO.Update(r.Context(), service, id, userID); err != nil {
		h.logger.Error("error calling update", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func (h Vehicle) delete(w http.ResponseWriter, r *http.Request, id, userID int64) {
	if err := h.serviceTypeDAO.Delete(r.Context(), id, userID); err != nil {
		h.logger.Error("error calling delete", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
