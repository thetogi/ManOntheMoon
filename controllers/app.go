package controllers

import (
	"ManOnTheMoonReviewService/controllers/response"
	"net/http"
	"strconv"
	"time"
)

type AppController struct {
	Controller
}

func (a *AppController) Home(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		w.WriteHeader(404)
		w.Write([]byte("Not Found !!!"))
		return
	}
	response.Write(w, response.Response{
		Code: http.StatusOK,
		Data: "Welcome to Man on the Moon homepage!",
	})
}

func (a *AppController) HealthCheck(w http.ResponseWriter, req *http.Request) {
	requestStartRaw := req.Header.Get("date")
	requestStart, _ := time.Parse(time.RFC1123, requestStartRaw)
	diff := time.Now().Sub(requestStart)
	diffString := strconv.FormatInt(diff.Milliseconds(), 10)

	response.Write(w, response.Response{
		Code: http.StatusOK,
		Data: "Man on the Moon Game Session Review service is running normally. Response time: " + diffString + " ms",
	})
}
