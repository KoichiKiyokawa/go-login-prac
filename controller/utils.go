package controller

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func respondJson(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, errors.WithStack(err).Error(), http.StatusInternalServerError)
	}
}
