package controllers

//Contains the methods that will handle the registered routes

import (
	"ManOnTheMoonReviewService/db"
	"encoding/json"
	"github.com/Pallinder/go-randomdata"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//Response contains HTTP response data
type Response struct {
	//Status indicates if a given request has succeeded, failed, etc.
	Status string

	//Message contains information regarding the outcome of a request
	Message string
}

//sendResponseAsJson returns a successful http status and provides response data as a JSON object
func sendResponseAsJson(w http.ResponseWriter, httpStatus int, responseData interface{}) {
	responseJson, err := json.Marshal(responseData)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(responseJson)
}

//*****GET Handlers*****//

func Home(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		w.WriteHeader(404)
		w.Write([]byte("Not Found !!!"))
		return
	}
	sendResponseAsJson(w, http.StatusOK, "Welcome to Man on the Moon homepage!")
}

func HealthCheck(w http.ResponseWriter, req *http.Request) {
	requestStartRaw := req.Header.Get("date")
	requestStart, _ := time.Parse(time.RFC1123, requestStartRaw)
	diff := time.Now().Sub(requestStart)
	diffString := strconv.FormatInt(diff.Milliseconds(), 10)
	sendResponseAsJson(w, http.StatusOK, "Man on the Moon Game Session Review service is running normally. Response time: "+diffString+" ms")
	//TODO this is how I generated feedback in readMe
}

//GetRatingRandomForTesting Simulates returning a rating. Nothing is retrieved from or committed to database.
func GetRatingRandomForTesting(w http.ResponseWriter, req *http.Request) {

	//Generate random session rating data
	sessionId := xid.New().String()
	playerId := xid.New().String()
	rand.Seed(time.Now().UnixNano())
	sessionRating := 1 + rand.Intn(5-1+1)
	sessionRatingComment := randomdata.Paragraph()

	//Trim random comment data to be less than 512 as that is the limit of the db comment field
	if len(sessionRatingComment) > 511 {
		sessionRatingComment = sessionRatingComment[0:511]
	}

	timeSubmitted := time.Now()
	data := db.SessionRating{PlayerId: playerId, SessionId: sessionId, Rating: sessionRating, Comment: sessionRatingComment, TimeSubmitted: timeSubmitted}

	sendResponseAsJson(w, http.StatusOK, data)
}

//GetPlayerByIdSql returns data of a single player by Id
func GetPlayerByIdSql(w http.ResponseWriter, req *http.Request) {

	//Get playerId from URL path
	params := mux.Vars(req)
	playerId := params["PlayerId"]

	//Retrieve player data by id
	playerData := db.SelectPlayerByPlayerId(playerId)

	//Check if player was retrieved and send response
	if playerData.PlayerId == "" {
		responseData := Response{Status: "PlayerNotFound", Message: "Could not find player using PlayerId: " + playerId}
		sendResponseAsJson(w, http.StatusBadRequest, responseData)
	} else {
		sendResponseAsJson(w, http.StatusOK, playerData)
	}
}

//GetAllPlayersSql returns all players from the players table
func GetAllPlayersSql(w http.ResponseWriter, req *http.Request) {

	//Retrieve all players from players table
	playersData := db.SelectAllPlayers()

	//Check if players were found then return response
	if len(playersData) == 0 {
		responseData := Response{Status: "PlayersNotFound", Message: "No players could be found."}
		sendResponseAsJson(w, http.StatusBadRequest, responseData)
	} else {
		//Return players data as a JSON response
		sendResponseAsJson(w, http.StatusOK, playersData)
	}
}

//GetSessionByIdSql returns data of a player's session by session and player Id
func GetSessionByIdSql(w http.ResponseWriter, req *http.Request) {

	//Get sessionId from URL path
	params := mux.Vars(req)
	sessionId := params["SessionId"]

	//Retrieve session data by Id
	sessionData := db.SelectSessionbyId(sessionId)

	//Check if session was retrieved and send response
	if sessionData.SessionId == "" {
		responseData := Response{Status: "SessionNotFound", Message: "Could not find session using SessionId: " + sessionId}
		sendResponseAsJson(w, http.StatusBadRequest, responseData)
	} else {
		sendResponseAsJson(w, http.StatusOK, sessionData)
	}
}

//GetAllSessionsSql returns all sessions by players from the sessions table
func GetAllSessionsSql(w http.ResponseWriter, req *http.Request) {

	//Retrieve all sessions from sessions table
	sessionsData := db.SelectAllSessions()
	sendResponseAsJson(w, http.StatusOK, sessionsData)

	//Check if sessions were found then return response
	if len(sessionsData) == 0 {
		responseData := Response{Status: "SessionsNotFound", Message: "No sessions could be found."}
		sendResponseAsJson(w, http.StatusBadRequest, responseData)
	} else {
		//Return players data as a JSON response
		sendResponseAsJson(w, http.StatusOK, sessionsData)
	}
}

//GetSessionRatingBySessionIdSql returns a player's session rating by their Ids
func GetSessionRatingBySessionIdSql(w http.ResponseWriter, req *http.Request) {

	//Get session from URL pathand player from URL path
	params := mux.Vars(req)
	sessionId := params["SessionId"]

	//Get playerId from URL query parameters
	playerId := req.URL.Query().Get("PlayerId")

	//Session Rating by session and player
	sessionRatingData := db.SelectSessionRating(sessionId, playerId)

	//Check if session rating was retrieved and send response
	if sessionRatingData.PlayerId == "" {
		responseData := Response{Status: "SessionRatingNotFound", Message: "Could not find rating by player for session using PlayerId: " + playerId + "and SessionId " + sessionId}
		sendResponseAsJson(w, http.StatusBadRequest, responseData)
	} else {
		sendResponseAsJson(w, http.StatusOK, sessionRatingData)
	}
}

//GetAllSessionRatingsSql returns all ratings by players for their sessions from the SessionRatings table. Optional filters can be provided for returning ratings.
func GetAllSessionRatingsSql(w http.ResponseWriter, req *http.Request) {
	//Get rating to filter by that value if provided
	rating := req.URL.Query().Get("Rating")

	//Get encoded rating filter operand if provided
	ratingFilterEnc := req.URL.Query().Get("Filter")

	//Get recent option if provided
	recentFlag := req.URL.Query().Get("Recent")

	//Validate recent flag
	if recentFlag != "" && recentFlag != "0" && recentFlag != "1" {
		responseData := Response{Status: "InvalidFlag", Message: "Recent parameter can only be a 0 or 1"}
		sendResponseAsJson(w, http.StatusBadRequest, responseData)
		return
	}

	//Validate rating filter
	var ratingFilter string
	var err error
	if ratingFilterEnc != "" {
		ratingFilter, err = url.QueryUnescape(ratingFilterEnc)
		if err != nil {
			responseData := Response{Status: "RatingFilterError", Message: err.Error()}
			sendResponseAsJson(w, http.StatusBadRequest, responseData)
			return
		}
		switch ratingFilter {
		case "<":
		case ">":
		case ">=":
		case "<=":
		default:
			responseData := Response{Status: "InvalidRatingFilter", Message: "Incorrect rating filter provided. Rating filter must be one of the following: <,<=,>,>="}
			sendResponseAsJson(w, http.StatusBadRequest, responseData)
			return
		}
	}

	if rating == "" && ratingFilter != "" {
		responseData := Response{Status: "NoRatingProvided", Message: "Rating was not provided with Filter."}
		sendResponseAsJson(w, http.StatusBadRequest, responseData)
	} else {

		//Convert rating as int
		ratingAsInt, _ := strconv.Atoi(rating)
		recentFlagAsBool, _ := strconv.ParseBool(recentFlag)

		sessionRatings := db.SelectAllSessionRatings(ratingAsInt, ratingFilter, recentFlagAsBool)

		//Check if session ratings were found then return response
		if len(sessionRatings) == 0 {
			responseData := Response{Status: "NoRatings", Message: "No ratings were found."}
			sendResponseAsJson(w, http.StatusBadRequest, responseData)
		} else {
			sendResponseAsJson(w, http.StatusOK, sessionRatings)
		}
	}
}

//*****POST Handlers*****//

func PostPlayerCreateSql(w http.ResponseWriter, req *http.Request) {

	//Generate new PlayerId
	playerID := xid.New().String()

	//Generate random Player name
	playerName := randomdata.FullName(randomdata.RandomGender)

	//Track Player being registered as current date and time
	PlayerTimeRegistered := time.Now()

	//Insert new Player into database
	ok, err := db.InsertNewPlayer(playerID, playerName, PlayerTimeRegistered)

	//If there is an error inserting, handle it
	if err != nil {
		panic(err)
	}

	//Check if insert was successful and send response
	var responseData Response
	var httpStatus int
	if ok == true {
		responseData = Response{Status: "OK", Message: "New Player Successfully created. ID: " + playerID}
		httpStatus = http.StatusOK
	} else {
		responseData = Response{Status: "FAILED", Message: "New Player was not created. ID: " + playerID}
		httpStatus = http.StatusBadRequest
	}

	//Send response as JSON
	sendResponseAsJson(w, httpStatus, responseData)
}

func PostSessionCreateSql(w http.ResponseWriter, req *http.Request) {

	//Generate new session id
	sessionId := xid.New().String()

	//Get parameters from URL query
	playerID := req.URL.Query().Get("PlayerId")
	timeSessionEnd, _ := time.Parse(time.RFC822, req.URL.Query().Get("TimeSessionEnd"))

	//Insert new Player into database
	ok, err := db.InsertNewSession(sessionId, playerID, timeSessionEnd)

	//If there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	//Check if insert was successful and send response
	var responseData Response
	var httpStatus int
	if ok == true {
		responseData = Response{Status: "OK", Message: "New Session Successfully created. ID: " + sessionId}
		httpStatus = http.StatusOK
	} else {
		responseData = Response{Status: "FAILED", Message: "Unable to create new session. ID: " + sessionId}
		httpStatus = http.StatusBadRequest
	}

	sendResponseAsJson(w, httpStatus, responseData)
}

func PostSessionRatingCreateSql(w http.ResponseWriter, req *http.Request) {

	//Get SessionId from parameter string
	params := mux.Vars(req)
	sessionId := params["SessionId"]
	playerId := req.URL.Query().Get("PlayerId")
	rating := req.URL.Query().Get("Rating")
	comment := req.URL.Query().Get("Comment")
	var responseData Response
	var httpStatus int
	//Check and prevent player from submitting another rating for the session if one exists, otherwise insert new rating
	currentRating := db.SelectSessionRating(sessionId, playerId)
	if currentRating.IsEmpty() {
		ratingInt, err := strconv.Atoi(rating)
		if err != nil {
			responseData = Response{Status: "FAILED_INVALID_RATING", Message: err.Error()}
			sendResponseAsJson(w, http.StatusBadRequest, responseData)
			return
		}

		if ratingInt < 1 || ratingInt > 5 {
			responseData = Response{Status: "FAILED_INVALID_RATING_QTY", Message: "Rating submitted is not valid. Ratings must be between 1 and 5."}
			sendResponseAsJson(w, http.StatusBadRequest, responseData)
			return
		}

		timeSubmitted := time.Now()

		//Insert new Player into database
		ok, err := db.InsertNewSessionRating(sessionId, playerId, ratingInt, comment, timeSubmitted)

		if err != nil {
			panic(err)
		}

		//Check if insert was successful and send response
		if ok == true {
			responseData = Response{Status: "OK", Message: "Rating Successfully submitted for Session ID: " + sessionId + " rating: " + rating + comment}
			httpStatus = http.StatusOK
		} else {
			responseData = Response{Status: "FAILED", Message: "Rating was unable to be submitted for Session ID: " + sessionId + " rating: " + rating + comment}
			httpStatus = http.StatusBadRequest
		}
	} else {
		responseData = Response{Status: "FAILED_DUPLICATE", Message: "Player has already submitted a rating for the session. Cannot submit more than one rating for a session. Session: " + sessionId + " Player: " + playerId + " rating: " + strconv.Itoa(currentRating.Rating) + " Comment: " + currentRating.Comment}
		httpStatus = http.StatusBadRequest
	}

	//Send response as JSON
	sendResponseAsJson(w, httpStatus, responseData)
}
