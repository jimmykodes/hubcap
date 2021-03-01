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

type ServiceType struct {
	logger         *zap.Logger
	serviceTypeDAO dao.ServiceType
}

func NewServiceType(logger *zap.Logger, serviceTypeDAO dao.ServiceType) *ServiceType {
	localLogger := logger.With(zap.String("handler", "service type"))
	return &ServiceType{logger: localLogger, serviceTypeDAO: serviceTypeDAO}
}

func (h ServiceType) Detail(w http.ResponseWriter, r *http.Request) {
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

func (h ServiceType) List(w http.ResponseWriter, r *http.Request) {
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

func (h ServiceType) create(w http.ResponseWriter, r *http.Request, userID int64) {
	serviceType := new(dto.ServiceType)
	if err := json.NewDecoder(r.Body).Decode(serviceType); err != nil {
		h.logger.Error("error decoding json", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	serviceType.UserID = userID
	if err := h.serviceTypeDAO.Create(r.Context(), serviceType); err != nil {
		h.logger.Error("error calling create", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (h ServiceType) list(w http.ResponseWriter, r *http.Request, userID int64) {
	sf := dao.SearchFilters{}
	for key, values := range r.URL.Query() {
		var value interface{}
		value = values[0]
		// not allowing multiple values. might do later
		switch key {
		case "freq_miles", "freq_days":
			var err error
			value, err = strconv.Atoi(values[0])
			if err != nil {
				writeErrorResponse(w, h.logger, http.StatusBadRequest, fmt.Sprintf("bad value for %s", key))
				return
			}
		}
		sf[key] = value
	}
	objs, err := h.serviceTypeDAO.Select(r.Context(), sf, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			objs = []*dto.ServiceType{}
		} else {
			h.logger.Error("error calling Select", zap.Error(err))
			writeErrorResponse(w, h.logger, http.StatusInternalServerError, "")
			return
		}
	}
	if err := json.NewEncoder(w).Encode(objs); err != nil {
		h.logger.Error("error writing data", zap.Error(err))
		writeErrorResponse(w, h.logger, http.StatusInternalServerError, "")
		return
	}

}
func (h ServiceType) get(w http.ResponseWriter, r *http.Request, id, userID int64) {
	serviceType, err := h.serviceTypeDAO.Get(r.Context(), id, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		h.logger.Error("error calling get", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(serviceType); err != nil {
		h.logger.Error("error writing data", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func (h ServiceType) update(w http.ResponseWriter, r *http.Request, id, userID int64) {
	serviceType := new(dto.ServiceType)
	if err := json.NewDecoder(r.Body).Decode(serviceType); err != nil {
		h.logger.Error("error decoding json", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := h.serviceTypeDAO.Update(r.Context(), serviceType, id, userID); err != nil {
		h.logger.Error("error calling update", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func (h ServiceType) delete(w http.ResponseWriter, r *http.Request, id, userID int64) {
	if err := h.serviceTypeDAO.Delete(r.Context(), id, userID); err != nil {
		h.logger.Error("error calling delete", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
