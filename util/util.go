package util

import (
	"ManOnTheMoonReviewService/controllers/response"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func NewUUID() string {
	return uuid.New().String()
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func ParseRequestBody(w http.ResponseWriter, r *http.Request, data interface{}) error {
	defer r.Body.Close()

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields() // catch unwanted fields

	err := d.Decode(&data)
	if err != nil {
		// bad JSON or unrecognized json field
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "ParseRequestBody",
			Message: "Unexpected error parsing request body",
		})
		return err
	}

	return nil
}
