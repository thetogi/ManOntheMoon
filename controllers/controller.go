package controllers

import (
	"ManOnTheMoonReviewService/controllers/response"
	"net/http"
)

//Base controller
type Controller struct{}

func (c *Controller) NotFound(w http.ResponseWriter, r *http.Request) {
	response.Write(w, response.Response{
		Code:    http.StatusNotFound,
		Message: "Not found",
	})
}
