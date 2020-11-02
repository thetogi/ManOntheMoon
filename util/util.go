package util

import (
	"ManOnTheMoonReviewService/controllers/response"
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
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

	r.Body = http.MaxBytesReader(w, r.Body, 100001)
	_, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Message: "Request too large",
			Errors:  map[string]string{"Error: ": err.Error()},
		})
		return err
	}

	d := json.NewDecoder(r.Body)

	// catch unwanted fields
	d.DisallowUnknownFields()

	err = d.Decode(&data)
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
