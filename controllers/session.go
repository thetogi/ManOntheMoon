package controllers

import (
	"ManOnTheMoonReviewService/controllers/response"
	"ManOnTheMoonReviewService/db"
	"ManOnTheMoonReviewService/models"
	"ManOnTheMoonReviewService/util"
	"net/http"
	"time"
)

type SessionController struct {
	Controller
}

//GetSessionByIdSql returns data of a player's session by session and player Id
func (s *SessionController) GetSession(w http.ResponseWriter, req *http.Request) {

	var session models.Session

	err := util.ParseRequestBody(w, req, &session)
	if err != nil {
		return
	}
	if session.SessionId == "" {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetSession",
			Message: "SessionId cannot be blank",
			Errors:  map[string]string{"SessionId": session.SessionId},
		})
		return
	}

	if !util.IsValidUUID(session.SessionId) {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetSession",
			Message: "SessionId is not a valid id",
			Errors:  map[string]string{"SessionId": session.SessionId},
		})
		return
	}

	//Retrieve session data by Id
	sessionData := db.SelectSession(session.SessionId)

	//Check if session was retrieved and send response
	if sessionData.SessionId == "" {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetSession",
			Message: "Could not find session.",
			Errors:  map[string]string{"SessionId": session.SessionId},
		})
	} else {
		response.Write(w, response.Response{
			Code: http.StatusOK,
			Data: sessionData,
		})
	}
}

func (s *SessionController) CreateSession(w http.ResponseWriter, req *http.Request) {

	var player models.Session

	err := util.ParseRequestBody(w, req, &player)
	if err != nil {
		return
	}

	if player.PlayerId == "" {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreateSession",
			Message: "PlayerId cannot be blank",
			Errors:  map[string]string{"PlayerId": player.PlayerId},
		})
		return
	}

	if !util.IsValidUUID(player.PlayerId) {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetPlayer",
			Message: "PlayerId is not a valid id",
			Errors:  map[string]string{"PlayerId": player.PlayerId},
		})
		return
	}
	var session models.Session
	session.SessionId = util.NewUUID()
	session.TimeSessionEnd = time.Now()

	//Insert new Player into database
	ok, err := db.InsertNewSession(session.SessionId, player.PlayerId, session.TimeSessionEnd)

	//If there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	//Check if insert was successful and send response
	var responseData response.Response
	if ok == true {
		responseData = response.Response{
			Code: http.StatusOK,
			Data: struct {
				Message string
				Data    models.Session
			}{
				"New Session Successfully created",
				session,
			},
		}
	} else {
		responseData = response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetRating",
			Message: "Unable to create new session for player",
			Errors:  map[string]string{"Player Name": player.PlayerId},
		}
	}
	response.Write(w, responseData)
}

//GetAllSessions returns all sessions by players from the sessions table
func (s *SessionController) GetAllSessions(w http.ResponseWriter, req *http.Request) {

	//Retrieve all sessions from sessions table
	sessionsData := db.SelectAllSessions()

	//Check if sessions were found then return response
	if len(sessionsData) == 0 {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetAllSessions",
			Message: "No sessions could be found.",
			Errors:  nil,
		})
	} else {
		//Return players data as a JSON response
		response.Write(w, response.Response{
			Code: http.StatusOK,
			Data: sessionsData,
		})
	}
}
