package util

import (
	"ManOnTheMoonReviewService/models"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func NewUUID() string {
	return uuid.New().String()
}

func ParseRequestBody(w http.ResponseWriter, r *http.Request, data interface{}) error {
	defer r.Body.Close()

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields() // catch unwanted fields
	var player models.Player

	err := d.Decode(&player)
	if err != nil {
		// bad JSON or unrecognized json field
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	return nil
}
