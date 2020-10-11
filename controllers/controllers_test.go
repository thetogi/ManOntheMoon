package controllers

import (
	"ManOnTheMoonReviewService/db"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRatingRandom(t *testing.T) {

	req, err := http.NewRequest("GET", "/GameSession/Rating", nil)

	checkError(err, t)

	rr := httptest.NewRecorder()

	http.HandlerFunc(GetRatingRandomForTesting).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var sr db.SessionRating

	err = json.Unmarshal(rr.Body.Bytes(), &sr)
	if err != nil {
		panic(err)
	}

	if sr.SessionId == "" || sr.PlayerId == "" {
		t.Errorf("No sessionId or PlayerId was generated")
	}
}

func TestGetPlayerByIdSql(t *testing.T) {

	playerId := "btvs3qli7nd6216va0rg"
	req, err := http.NewRequest("GET", "/Player/"+playerId, nil)
	req = mux.SetURLVars(req, map[string]string{"PlayerId": "btvs3qli7nd6216va0rg"})
	checkError(err, t)

	rr := httptest.NewRecorder()

	http.HandlerFunc(GetPlayerByIdSql).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var p db.Player

	err = json.Unmarshal(rr.Body.Bytes(), &p)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, playerId, p.PlayerId, "PlayerId differs")

	if p.PlayerId == "" {
		t.Errorf("Error finding player Id.")
	}
}

func TestGetAllPlayersSql(t *testing.T) {

	req, err := http.NewRequest("GET", "/Players/", nil)
	checkError(err, t)

	rr := httptest.NewRecorder()

	http.HandlerFunc(GetAllPlayersSql).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var p []db.Player

	err = json.Unmarshal(rr.Body.Bytes(), &p)
	if err != nil {
		panic(err)
	}

	if len(p) == 0 {
		t.Errorf("Error retrieving players")
	}
}

func TestGetSessionByIdSql(t *testing.T) {

	sessionId := "btvtm3di7nd52o66pn70"
	req, err := http.NewRequest("GET", "/Session/"+sessionId, nil)
	req = mux.SetURLVars(req, map[string]string{"SessionId": "btvtm3di7nd52o66pn70"})
	checkError(err, t)

	rr := httptest.NewRecorder()

	http.HandlerFunc(GetSessionByIdSql).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var s db.Session

	err = json.Unmarshal(rr.Body.Bytes(), &s)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, sessionId, s.SessionId, "SessionId differs")

	if s.SessionId == "" {
		t.Errorf("Error finding session Id.")
	}
}

func TestGetAllSessionsSql(t *testing.T) {

	req, err := http.NewRequest("GET", "/Sessions/", nil)
	checkError(err, t)

	rr := httptest.NewRecorder()

	http.HandlerFunc(GetAllPlayersSql).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var p []db.Session

	err = json.Unmarshal(rr.Body.Bytes(), &p)
	if err != nil {
		panic(err)
	}

	if len(p) == 0 {
		t.Errorf("Error retrieving sessions")
	}
}

func TestGetSessionRatingBySessionIdSql(t *testing.T) {

	sessionId := "btvtm3di7nd52o66pn70"
	req, err := http.NewRequest("GET", "/Session/"+sessionId+"/Rating", nil)
	checkError(err, t)
	req = mux.SetURLVars(req, map[string]string{"SessionId": sessionId})

	q := req.URL.Query()
	q.Add("PlayerId", "btvs3qli7nd6216va0rg")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()

	http.HandlerFunc(GetSessionRatingBySessionIdSql).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var sr db.SessionRating

	err = json.Unmarshal(rr.Body.Bytes(), &sr)
	if err != nil {
		panic(err)
	}

	if sr.SessionId == "" || sr.PlayerId == "" {
		t.Errorf("No sessionId or PlayerId was generated")
	}
}

func TestGetAllSessionRatingsSql(t *testing.T) {

	req, err := http.NewRequest("GET", "/Session/Ratings/", nil)
	checkError(err, t)

	rr := httptest.NewRecorder()

	http.HandlerFunc(GetAllSessionRatingsSql).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var p []db.SessionRating

	err = json.Unmarshal(rr.Body.Bytes(), &p)
	if err != nil {
		panic(err)
	}

	if len(p) == 0 {
		t.Errorf("Error retrieving Ratings.")
	}
}

func TestPostPlayerCreateSql(t *testing.T) {

	req, err := http.NewRequest("POST", "/Player/Create", nil)
	checkError(err, t)

	rr := httptest.NewRecorder()

	http.HandlerFunc(PostPlayerCreateSql).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var response Response

	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		panic(err)
	}

	if response.Status == "" || response.Status != "OK" {
		//TODO handle blank message
		t.Errorf("Error creating new player: " + response.Message)
	}
}

func TestPostSessionCreateSql(t *testing.T) {

	req, err := http.NewRequest("POST", "/Session/Create", nil)
	checkError(err, t)

	q := req.URL.Query()
	q.Add("PlayerId", "btvs3qli7nd6216va0rg")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()

	http.HandlerFunc(PostSessionCreateSql).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var response Response

	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		panic(err)
	}

	if response.Status == "" || response.Status != "OK" {
		//TODO handle blank message
		t.Errorf("Error creating new session: " + response.Message)
	}
}

func TestPostSessionRatingCreateSql(t *testing.T) {

	sessionId := "btvtm3di7nd52o66pn70"
	req, err := http.NewRequest("POST", "/Session/"+sessionId+"/CreateRating", nil)
	checkError(err, t)

	q := req.URL.Query()
	q.Add("PlayerId", "btvs3qli7nd6216va0rg")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()

	http.HandlerFunc(PostSessionRatingCreateSql).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var response Response

	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		panic(err)
	}

	if (response.Status == "" || response.Status != "OK") && response.Status != "FAILED_DUPLICATE" {
		//TODO handle blank message
		t.Errorf("Error creating new session rating: " + response.Message)
	}
}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}
