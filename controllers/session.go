package controllers

import (
	"ManOnTheMoonReviewService/controllers/response"
	"ManOnTheMoonReviewService/db"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	"net/http"
	"time"
)

type SessionController struct {
	Controller
}

//GetSessionByIdSql returns data of a player's session by session and player Id
func (*SessionController) GetSession(w http.ResponseWriter, req *http.Request) {

	//Get sessionId from URL path
	params := mux.Vars(req)
	sessionId := params["SessionId"]

	//Retrieve session data by Id
	sessionData := db.SelectSession(sessionId)

	//Check if session was retrieved and send response
	if sessionData.SessionId == "" {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetSession",
			Message: "Could not find session using SessionId: " + sessionId,
		})
	} else {
		response.Write(w, response.Response{
			Code: http.StatusOK,
			Data: sessionData,
		})
	}
}

func (*SessionController) CreateSession(w http.ResponseWriter, req *http.Request) {

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
	var responseData response.Response
	if ok == true {
		responseData = response.Response{
			Code: http.StatusOK,
			Data: "New Session Successfully created. ID: " + sessionId,
		}
	} else {
		responseData = response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetRating",
			Message: "Unable to create new session. ID: " + sessionId,
		}
	}
	response.Write(w, responseData)
}

//GetAllSessions returns all sessions by players from the sessions table
func (*SessionController) GetAllSessions(w http.ResponseWriter, req *http.Request) {

	//Retrieve all sessions from sessions table
	sessionsData := db.SelectAllSessions()

	//Check if sessions were found then return response
	if len(sessionsData) == 0 {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetAllSessions",
			Message: "No sessions could be found.",
		})
	} else {
		//Return players data as a JSON response
		response.Write(w, response.Response{
			Code: http.StatusOK,
			Data: sessionsData,
		})
	}
}
