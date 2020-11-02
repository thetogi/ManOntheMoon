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

//Parses the body of a request, validates body size to prevent large requests from accidentally or maliciously stopping server
func ParseRequestBody(w http.ResponseWriter, r *http.Request, data interface{}) error {
	defer r.Body.Close()

	d := json.NewDecoder(r.Body)

	// catch unwanted fields
	d.DisallowUnknownFields()

	err := d.Decode(&data)
	if err != nil {
		// bad JSON or unrecognized field in request
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "ParseRequestBody",
			Message: "Unexpected data in request body",
		})
		return err
	}

	return nil
}
