package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/guilherme8/go-api/models"
)

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}
