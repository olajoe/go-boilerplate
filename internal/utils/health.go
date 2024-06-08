package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).
		Encode(map[string]interface{}{
			"message": "OK",
		}); err != nil {
		slog.Error(err.Error())
	}
}
