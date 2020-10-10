package controllers

import (
	"ManOnTheMoonReviewService/db"
	"ManOnTheMoonReviewService/util"
	"encoding/json"
	"github.com/Pallinder/go-randomdata"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Players struct {
	Data []db.Player
}

type Sessions struct {
	Data []db.Session
}

type SessionRatings struct {
	Data []db.SessionRating
}

//Http Response Structs
type ResponsePostSessionRating struct {
	Status  string
	Message string
}

type ResponsePost struct {
	Status  string
	Message string
}

type ResponseGet struct {
	Status  string
	Message string
}

func sendResponseAsJson(w http.ResponseWriter, data interface{}) {
	responsePost, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responsePost)
}

func ShowEnvToggle(w http.ResponseWriter, req *http.Request) {

	files := []string{
		"./ui/html/envToggle.layout.tmpl",
		"./ui/html/envToggle.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "Internal Server Error", 500)

	}
	if util.EnvShowStatus() == false {
		util.ToggleEnvShow()
	}

	err = ts.Execute(w, util.EnvData())
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}

func Home(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		w.WriteHeader(404)
		w.Write([]byte("Not Found !!!"))
		return
	}
	newSessionId := xid.New()
	//Store session ID in DB

	// Redirect to new session
	http.Redirect(w, req, "http://localhost:8080/GameSession/"+newSessionId.String(), http.StatusSeeOther)

	return
	//files := []string{
	//	"./ui/html/home.page.tmpl",
	//	"./ui/html/base.layout.tmpl",
	//	"./ui/html/footer.partial.tmpl",
	//}
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	log.Printf(err.Error())
	//	http.Error(w, "Internal Server Error", 500)
	//
	//}
	//
	//err = ts.Execute(w, util.EnvData())
	//if err != nil {
	//	log.Printf(err.Error())
	//	http.Error(w, "Internal Server Error", 500)
	//}
}

func GetRatingRandom(w http.ResponseWriter, req *http.Request) {

	playerID := xid.New().String()
	session := xid.New().String()
	rand.Seed(time.Now().UnixNano())
	sessionRating := 1 + rand.Intn(5-1+1)
	sessionRatingComment := randomdata.Paragraph()

	if len(sessionRatingComment) > 511 {
		sessionRatingComment = sessionRatingComment[0:511]
	}
	TimeSubmitted := time.Now()
	data := db.SessionRating{PlayerId: playerID, SessionId: session, Rating: sessionRating, Comment: sessionRatingComment, TimeSubmitted: TimeSubmitted}

	sendResponseAsJson(w, data)
}

func PostGameCreateSql(w http.ResponseWriter, req *http.Request) {

	gameId := xid.New().String()
	gameName := "ManOnTheMoon"
	gameVersion := "1.0"
	sqlStatement := "INSERT INTO games (`GameId`,`Name`,`Version`) VALUES ( ?, ?, ?)"
	insert, err := db.Db.Query(sqlStatement, gameId, gameName, gameVersion)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	data := ResponsePost{Status: "OK", Message: "New GAme Successfully created. ID: " + gameId}

	sendResponseAsJson(w, data)
}

//Finds and returns data about a Game by GameId
func GetGameByIdSql(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	gameId := params["GameId"]

	//Check for blank data
	if gameId == "" {
		//TODO Throw Error
	}

	//Retrieve game data
	gameData := db.SelectGameById(gameId)

	//Return game data as a JSON response
	sendResponseAsJson(w, gameData)
}

func PostPlayerCreateSql(w http.ResponseWriter, req *http.Request) {

	//Generate new PlayerId
	playerID := xid.New().String()

	//Generate random Player name
	playerName := randomdata.FullName(randomdata.RandomGender)

	//Track Player being registered as current date and time
	PlayerTimeRegistered := time.Now()

	//Insert new Player into database
	ok, err := db.InsertNewPlayer(playerID, playerName, PlayerTimeRegistered)

	if err != nil {
		panic(err)
	}

	//Build response based on successful player insert
	var responseData ResponsePost

	if ok == true {
		responseData = ResponsePost{Status: "OK", Message: "New Player Successfully created. ID: " + playerID}
	} else {
		responseData = ResponsePost{Status: "FAILED", Message: "New Player was not created. ID: " + playerID}
	}

	//Send response as JSON
	sendResponseAsJson(w, responseData)
}

func GetPlayerByIdSql(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	playerId := params["PlayerId"]

	//Check for blank data
	if playerId == "" {
		//TODO Throw Error
	}

	//Retrieve player data
	playerData := db.SelectPlayerByPlayerId(playerId)

	//Return player data as a JSON response
	sendResponseAsJson(w, playerData)
}

func GetAllPlayersSql(w http.ResponseWriter, req *http.Request) {

	//Retrieve all players from table
	playersData := db.SelectAllPlayers()

	//Return players data as a JSON response
	sendResponseAsJson(w, playersData)
}

func PostSessionCreateSql(w http.ResponseWriter, req *http.Request) {

	//Generate new session id
	sessionId := xid.New().String()

	//Get parameters from URL query
	playerID := req.URL.Query().Get("PlayerId")
	gameId := req.URL.Query().Get("GameId")
	timeSessionEnd, _ := time.Parse(time.RFC822, req.URL.Query().Get("TimeSessionEnd"))

	//Insert new Player into database
	ok, err := db.InsertNewSession(sessionId, gameId, playerID, timeSessionEnd)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	//Build response based on successful player insert
	var responseData ResponsePost

	if ok == true {
		responseData = ResponsePost{Status: "OK", Message: "New Session Successfully created. ID: " + sessionId}
	} else {
		responseData = ResponsePost{Status: "FAILED", Message: "Unable to create new session. ID: " + sessionId}
	}

	sendResponseAsJson(w, responseData)
}

func GetSessionByIdSql(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	sessionId := params["SessionId"]

	sessionData := db.SelectSessionbyId(sessionId)

	sendResponseAsJson(w, sessionData)
}

func GetAllSessionsSql(w http.ResponseWriter, req *http.Request) {

	sessionsData := db.SelectAllSessions()
	sendResponseAsJson(w, sessionsData)
}

func PostSessionRatingCreateSql(w http.ResponseWriter, req *http.Request) {
	//Get SessionId from parameter string
	params := mux.Vars(req)
	sessionId := params["SessionId"]
	playerId := params["PlayerId"]

	//Check and prevent player from submitting another rating for the session if one exists

	req.ParseForm()
	rating, _ := strconv.Atoi(req.FormValue("Rating"))
	comment := req.Form.Get("Comment")
	timeSubmitted := time.Now()

	//Insert new Player into database
	ok, err := db.InsertNewSessionRating(sessionId, playerId, rating, comment, timeSubmitted)

	if err != nil {
		panic(err)
	}

	//Build response based on successful player insert
	var responseData ResponsePost

	if ok == true {
		responseData = ResponsePost{Status: "OK", Message: "Rating Successfully submitted for Session ID: " + sessionId + " rating: " + strconv.Itoa(rating) + comment}
	} else {
		responseData = ResponsePost{Status: "FAILED", Message: "Rating was unable to be submitted for Session ID: " + sessionId + " rating: " + strconv.Itoa(rating) + comment}
	}

	//Send response as JSON
	sendResponseAsJson(w, responseData)
}

func GetSessionRatingBySessionIdSql(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	sessionId := params["SessionId"]
	playerId := params["PlayerId"]

	sessionRatingData := db.SelectSessionRatingBySessionId(sessionId, playerId)

	sendResponseAsJson(w, sessionRatingData)
}

func GetAllSessionRatingsSql(w http.ResponseWriter, req *http.Request) {
	rating, _ := strconv.Atoi(req.URL.Query().Get("Rating"))
	//TODO handle parsing errors
	ratingFilter := req.URL.Query().Get("Filter")

	//ratingFilterOp := req.URL.Query().Get("Op")

	sessionRatings := db.SelectAllSessionRatings(rating, ratingFilter)
	sendResponseAsJson(w, sessionRatings)
}
