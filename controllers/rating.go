package controllers

import (
	"ManOnTheMoonReviewService/controllers/response"
	"ManOnTheMoonReviewService/db"
	seed "ManOnTheMoonReviewService/db/seed/seeder"
	"ManOnTheMoonReviewService/models"
	"ManOnTheMoonReviewService/util"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type RatingController struct {
	Controller
}

//GetRating returns a player's session rating by their Ids
func (r *RatingController) GetRating(w http.ResponseWriter, req *http.Request) {

	//Get session from URL path and player from URL path
	params := mux.Vars(req)
	sessionId := params["SessionId"]

	//Get playerId from URL query parameters
	playerId := req.URL.Query().Get("PlayerId")

	//Session Rating by session and player
	ratingData := db.SelectRating(sessionId, playerId)

	//Check if session rating was retrieved and send response
	if ratingData.PlayerId == "" {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetRating",
			Message: "Could not find rating by player for session using PlayerId: " + playerId + "and SessionId " + sessionId,
		})
	} else {
		response.Write(w, response.Response{
			Code: http.StatusOK,
			Data: ratingData,
		})
	}
}

//GetRatings returns all ratings by players for their sessions from the Ratings table. Optional filters can be provided for returning ratings.
func (r *RatingController) GetRatings(w http.ResponseWriter, req *http.Request) {
	//Get rating to filter by that value if provided
	rating := req.URL.Query().Get("Rating")

	//Get encoded rating filter operand if provided
	ratingFilterEnc := req.URL.Query().Get("Filter")

	//Get recent option if provided
	recentFlag := req.URL.Query().Get("Recent")

	//Validate recent flag
	if recentFlag != "" && recentFlag != "0" && recentFlag != "1" {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetAllSessions",
			Message: "Recent parameter can only be a 0 or 1",
		})
		return
	}

	//Validate rating filter
	var ratingFilter string
	var err error
	if ratingFilterEnc != "" {
		ratingFilter, err = url.QueryUnescape(ratingFilterEnc)
		if err != nil {
			response.Write(w, response.Response{
				Code:    http.StatusBadRequest,
				Action:  "GetAllSessions",
				Message: "Error handling filter parameter. Log: " + err.Error(),
			})
			return
		}
		switch ratingFilter {
		case "<":
		case ">":
		case ">=":
		case "<=":
		default:
			response.Write(w, response.Response{
				Code:    http.StatusBadRequest,
				Action:  "GetAllSessions",
				Message: "Incorrect rating filter provided. Rating filter must be one of the following: <,<=,>,>=",
			})
			return
		}
	}

	if rating == "" && ratingFilter != "" {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetAllSessions",
			Message: "Rating was not provided with Filter.",
		})
	} else {

		//Convert rating as int
		ratingAsInt, _ := strconv.Atoi(rating)
		recentFlagAsBool, _ := strconv.ParseBool(recentFlag)

		ratings := db.SelectAllRatings(ratingAsInt, ratingFilter, recentFlagAsBool)

		//Check if session ratings were found then return response
		if len(ratings) == 0 {
			response.Write(w, response.Response{
				Code:    http.StatusBadRequest,
				Action:  "GetAllSessions",
				Message: "No ratings were found.",
			})
		} else {
			response.Write(w, response.Response{
				Code: http.StatusOK,
				Data: ratings,
			})
		}
	}
}

func (r *RatingController) CreateRating(w http.ResponseWriter, req *http.Request) {

	var rating models.Rating
	err := util.ParseRequestBody(w, req, &rating)
	if err != nil {
		return
	}

	if rating.PlayerId == "" {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreateRating",
			Message: "PlayerId cannot be blank",
			Errors:  map[string]string{"PlayerId": rating.PlayerId},
		})
		return
	}

	if !util.IsValidUUID(rating.PlayerId) {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreateRating",
			Message: "PlayerId is not a valid id",
			Errors:  map[string]string{"PlayerId": rating.PlayerId},
		})
		return
	}

	if rating.SessionId == "" {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreateRating",
			Message: "SessionId cannot be blank",
			Errors:  map[string]string{"SessionId": rating.SessionId},
		})
		return
	}

	if !util.IsValidUUID(rating.SessionId) {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreateRating",
			Message: "SessionId is not a valid id",
			Errors:  map[string]string{"SessionId": rating.SessionId},
		})
		return
	}

	if rating.Rating == 0 {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreateRating",
			Message: "Rating cannot be 0",
			Errors:  map[string]string{"Rating": "0"},
		})
		return
	}

	var responseData response.Response

	//Check and prevent player from submitting another rating for the session if one exists, otherwise insert new rating
	currentRating := db.SelectRating(rating.SessionId, rating.PlayerId)
	if !currentRating.IsEmpty() {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreateRating",
			Message: "Player has already submitted a rating for the session. Cannot submit more than one rating for a session",
			Errors: map[string]string{
				"PlayerId":       rating.PlayerId,
				"SessionId":      rating.SessionId,
				"CurrentRating":  strconv.Itoa(currentRating.Rating),
				"CurrentComment": currentRating.Comment,
			},
		})
		return
	}

	if rating.Rating < 1 || rating.Rating > 5 {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreateRating",
			Message: "Rating submitted is not valid. Ratings must be between 1 and 5.",
		})
		return
	}

	timeSubmitted := time.Now()

	//Insert new Player into database
	ok, err := db.InsertNewRating(
		rating.SessionId,
		rating.PlayerId,
		rating.Rating,
		rating.Comment,
		timeSubmitted,
	)

	if err != nil {
		panic(err)
	}

	//Check if insert was successful and send response
	if ok == true {
		responseData = response.Response{
			Code: http.StatusOK,
			Data: struct {
				Message string
				Data    models.Rating
			}{
				"Rating Successfully submitted for Session",
				rating,
			},
		}
	} else {
		responseData = response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreateRating",
			Message: "Failed to create rating for session",
			Errors: map[string]string{
				"PlayerId":  rating.PlayerId,
				"SessionId": rating.SessionId,
				"Rating":    strconv.Itoa(rating.Rating),
				"Comment":   rating.Comment,
			},
		}
	}
	response.Write(w, responseData)
}

//GetRandomRating Simulates returning a rating. Nothing is retrieved from or committed to database.
func (r *RatingController) GetRandomRating(w http.ResponseWriter, req *http.Request) {
	response.Write(w, response.Response{
		Code: http.StatusOK,
		Data: seed.MockRatingData(),
	})
}
