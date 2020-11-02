package controllers

import (
	"ManOnTheMoonReviewService/controllers/response"
	seed "ManOnTheMoonReviewService/db/seed/seeder"
	"ManOnTheMoonReviewService/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

var ExistingPlayerId string
var ExistingSessionId string

func init() {
	ExistingPlayerId = "bu2hhrti7nd0md7faom0"
	ExistingSessionId = "bu2hhrti7nd0md7faoog"
}

func TestRatingController_GetRandomRating(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/v1/Random/Rating", nil)
	rr := httptest.NewRecorder()
	rc := RatingController{}

	http.HandlerFunc(rc.GetRandomRating).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var r response.Response

	err = json.Unmarshal(rr.Body.Bytes(), &r)
	if err != nil {
		panic(err)
	}

	receivedDTO := response.Response{}

	err = json.Unmarshal(rr.Body.Bytes(), &receivedDTO)
	if err != nil {
		panic(err)
	}

	rating := receivedDTO.Data.(map[string]interface{})
	if rating["SessionId"] == "" || rating["PlayerId"] == "" {
		t.Errorf("No sessionId or PlayerId was generated")
	}
}

func TestGetPlayerByIdSql(t *testing.T) {
	bodyRaw := "{\"playerId\":\"2675e3f6-22db-4253-8d5c-eb7d8cfa333c\"}"

	b, err := json.Marshal(bodyRaw)
	if err != nil {
		fmt.Println("error:", err)
	}
	r := bytes.NewReader(b)
	req, err := http.NewRequest("GET", "/api/v1/Player/", r)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//Use SetURLVars for tests so that the handler will correctly retrieve the URL path parameters
	req = mux.SetURLVars(req, map[string]string{"PlayerId": ExistingPlayerId})
	rr := httptest.NewRecorder()
	pc := PlayerController{}
	http.HandlerFunc(pc.GetPlayer).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var p models.Player

	err = json.Unmarshal(rr.Body.Bytes(), &p)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, ExistingPlayerId, p.PlayerId, "PlayerId differs")

	if p.PlayerId == "" {
		t.Errorf("Error finding player Id.")
	}
}

func TestGetAllPlayersSql(t *testing.T) {

	req, err := http.NewRequest("GET", "/Players/", nil)

	rr := httptest.NewRecorder()
	pc := PlayerController{}
	http.HandlerFunc(pc.GetAllPlayers).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var p models.Players

	err = json.Unmarshal(rr.Body.Bytes(), &p)
	if err != nil {
		panic(err)
	}

	if len(p) == 0 {
		t.Errorf("Error retrieving players")
	}
}

func TestGetSessionByIdSql(t *testing.T) {

	req, err := http.NewRequest("GET", "/Session/"+ExistingSessionId, nil)

	//Use SetURLVars for tests so that the handler will correctly retrieve the URL path parameters
	req = mux.SetURLVars(req, map[string]string{"SessionId": ExistingSessionId})
	rr := httptest.NewRecorder()
	sc := SessionController{}
	http.HandlerFunc(sc.GetSession).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var s models.Session

	err = json.Unmarshal(rr.Body.Bytes(), &s)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, ExistingSessionId, s.SessionId, "SessionId differs")

	if s.SessionId == "" {
		t.Errorf("Error finding session Id.")
	}
}

func TestGetAllSessionsSql(t *testing.T) {

	req, err := http.NewRequest("GET", "/Sessions/", nil)

	rr := httptest.NewRecorder()
	sc := SessionController{}
	http.HandlerFunc(sc.GetAllSessions).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var p models.Sessions

	err = json.Unmarshal(rr.Body.Bytes(), &p)
	if err != nil {
		panic(err)
	}

	if len(p) == 0 {
		t.Errorf("Error retrieving sessions")
	}
}

func TestGetSessionRatingBySessionIdSql(t *testing.T) {

	req, err := http.NewRequest("GET", "/Session/"+ExistingSessionId+"/Rating", nil)

	//Use SetURLVars for tests so that the handler will correctly retrieve the URL path parameters
	req = mux.SetURLVars(req, map[string]string{"SessionId": ExistingSessionId})

	//Set query parameters
	q := req.URL.Query()
	q.Add("PlayerId", ExistingPlayerId)
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	rc := RatingController{}
	http.HandlerFunc(rc.GetRating).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var r models.Rating

	err = json.Unmarshal(rr.Body.Bytes(), &r)
	if err != nil {
		panic(err)
	}

	if r.SessionId == "" || r.PlayerId == "" {
		t.Errorf("No sessionId or PlayerId was generated")
	}
}

func TestGetAllRatingsSql(t *testing.T) {

	req, err := http.NewRequest("GET", "/Session/Ratings/", nil)

	rr := httptest.NewRecorder()
	rc := RatingController{}
	http.HandlerFunc(rc.GetRatings).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var r models.Ratings

	err = json.Unmarshal(rr.Body.Bytes(), &r.Data)
	if err != nil {
		panic(err)
	}

	if len(r.Data) == 0 {
		t.Errorf("Error retrieving Ratings.")
	}
}

func TestPostPlayerCreateSql(t *testing.T) {

	req, err := http.NewRequest("POST", "/Player/Create", nil)

	rr := httptest.NewRecorder()
	pc := PlayerController{}
	http.HandlerFunc(pc.CreatePlayer).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var response response.Response

	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		panic(err)
	}

	if response.Code != http.StatusOK {
		//TODO handle blank message
		t.Errorf("Error creating new player: " + response.Message + " StatusCode: " + strconv.Itoa(response.Code))
	}
}

func TestPostSessionCreateSql(t *testing.T) {

	req, err := http.NewRequest("POST", "/Session/Create", nil)

	q := req.URL.Query()
	q.Add("PlayerId", ExistingPlayerId)
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	sc := SessionController{}
	http.HandlerFunc(sc.CreateSession).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var response response.Response

	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		panic(err)
	}

	if response.Code != http.StatusOK {
		//TODO handle blank message
		t.Errorf("Error creating new session: " + response.Message + " StatusCode: " + strconv.Itoa(response.Code))
	}
}

func TestPostSessionRatingCreateSql(t *testing.T) {

	newRating := seed.MockRatingData()
	b, err := json.Marshal(newRating)
	if err != nil {
		fmt.Println("error:", err)
	}
	req, err := http.NewRequest("POST", "/api/v1/Rating", strings.NewReader(string(b)))
	rr := httptest.NewRecorder()
	rc := RatingController{}
	http.HandlerFunc(rc.CreateRating).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d. Got %d", http.StatusOK, status)
	}

	var response response.Response

	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		panic(err)
	}

	if response.Code != http.StatusOK {
		{
			//TODO handle blank message
			t.Errorf("Error creating new session rating: " + response.Message + " StatusCode: " + strconv.Itoa(response.Code))
		}
	}
}
