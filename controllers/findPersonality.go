package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/guilherme8/go-api/models"
)

func FindPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personasilty := range models.Personalities {
		if strconv.Itoa(personasilty.Id) == id {
			json.NewEncoder(w).Encode(personasilty)
		}
	}
}
