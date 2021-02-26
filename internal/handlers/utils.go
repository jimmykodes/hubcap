package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type ContextKey int

const (
	userIDKey ContextKey = iota
	userKey
)

func writeErrorResponse(w http.ResponseWriter, logger *zap.Logger, header int, msg string) {
	w.WriteHeader(header)
	if msg != "" {
		if err := json.NewEncoder(w).Encode(map[string]string{"error": msg}); err != nil {
			logger.Error("error writing message", zap.Error(err), zap.String("message", msg))
		}
	}
}
