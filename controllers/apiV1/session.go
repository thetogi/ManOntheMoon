package apiV1

import (
	"ManOnTheMoonReviewService/controllers"
	"ManOnTheMoonReviewService/controllers/response"
	"ManOnTheMoonReviewService/models"
	"ManOnTheMoonReviewService/util"
	"net/http"
)

type SessionController struct {
	controllers.Controller
	Session  models.Session
	Sessions models.Sessions
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
	s.Session.Retrieve(&session)

	//Check if session was retrieved and send response
	if session.SessionId == "" {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetSession",
			Message: "Could not find session.",
			Errors:  map[string]string{"SessionId": session.SessionId},
		})
	} else {
		response.Write(w, response.Response{
			Code: http.StatusOK,
			Data: session,
		})
	}
}

func (s *SessionController) CreateSession(w http.ResponseWriter, req *http.Request) {

	var session models.Session

	err := util.ParseRequestBody(w, req, &session)
	if err != nil {
		return
	}

	if session.PlayerId == "" {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreateSession",
			Message: "PlayerId cannot be blank",
			Errors:  map[string]string{"PlayerId": session.PlayerId},
		})
		return
	}

	if !util.IsValidUUID(session.PlayerId) {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetPlayer",
			Message: "PlayerId is not a valid id",
			Errors:  map[string]string{"PlayerId": session.PlayerId},
		})
		return
	}

	ok, err := s.Session.Create(&session)

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
			Message: "Unable to create new session",
			Errors:  map[string]string{"Player Id": session.PlayerId},
		}
	}
	response.Write(w, responseData)
}

//GetAllSessions returns all sessions by players from the sessions table
func (s *SessionController) GetAllSessions(w http.ResponseWriter, req *http.Request) {

	//Retrieve all sessions from sessions table
	var sessions models.Sessions
	s.Session.RetrieveAll(&sessions)

	//Check if sessions were found then return response
	if len(sessions) == 0 {
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
			Data: sessions,
		})
	}
}
