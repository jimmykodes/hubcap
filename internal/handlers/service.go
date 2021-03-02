package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/jimmykodes/vehicle_maintenance/internal/dao"
	"github.com/jimmykodes/vehicle_maintenance/internal/dto"
)

type Service struct {
	logger     *zap.Logger
	serviceDAO dao.Service
}

func NewService(logger *zap.Logger, serviceDAO dao.Service) *Service {
	localLogger := logger.With(zap.String("handler", "service"))
	return &Service{logger: localLogger, serviceDAO: serviceDAO}
}

func (h Service) Detail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		h.logger.Error("error parsing service ID", zap.Error(err))
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

func (h Service) List(w http.ResponseWriter, r *http.Request) {
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

func (h Service) create(w http.ResponseWriter, r *http.Request, userID int64) {
	service := new(dto.Service)
	if err := json.NewDecoder(r.Body).Decode(service); err != nil {
		h.logger.Error("error decoding service json", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	service.UserID = userID
	if err := h.serviceDAO.Create(r.Context(), service); err != nil {
		h.logger.Error("error calling create", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (h Service) list(w http.ResponseWriter, r *http.Request, userID int64) {
	sf := dao.SearchFilters{}
	for key, values := range r.URL.Query() {
		var value interface{}
		value = values[0]
		// not allowing multiple values. might do later
		switch key {
		case "service_type_id", "vehicle_id":
			var err error
			value, err = strconv.Atoi(values[0])
			if err != nil {
				writeErrorResponse(w, h.logger, http.StatusBadRequest, fmt.Sprintf("bad value for %s", key))
				return
			}
		}
		sf[key] = value
	}
	objs, err := h.serviceDAO.Select(r.Context(), sf, userID)
	if err != nil {
		h.logger.Error("error calling Select", zap.Error(err))
		writeErrorResponse(w, h.logger, http.StatusInternalServerError, "")
		return
	}
	if objs == nil {
		objs = []*dto.Service{}
	}
	if err := json.NewEncoder(w).Encode(objs); err != nil {
		h.logger.Error("error writing data", zap.Error(err))
		writeErrorResponse(w, h.logger, http.StatusInternalServerError, "")
		return
	}
}
func (h Service) get(w http.ResponseWriter, r *http.Request, id, userID int64) {
	service, err := h.serviceDAO.Get(r.Context(), id, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		h.logger.Error("error getting service", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(service); err != nil {
		h.logger.Error("error writing service data", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func (h Service) update(w http.ResponseWriter, r *http.Request, id, userID int64) {
	service := new(dto.Service)
	if err := json.NewDecoder(r.Body).Decode(service); err != nil {
		h.logger.Error("error decoding service json", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := h.serviceDAO.Update(r.Context(), service, id, userID); err != nil {
		h.logger.Error("error updating service", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func (h Service) delete(w http.ResponseWriter, r *http.Request, id, userID int64) {
	if err := h.serviceDAO.Delete(r.Context(), id, userID); err != nil {
		h.logger.Error("error deleting service", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
