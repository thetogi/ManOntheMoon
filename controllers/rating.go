package controllers

import (
	"ManOnTheMoonReviewService/controllers/response"
	"ManOnTheMoonReviewService/db"
	"ManOnTheMoonReviewService/models"
	"github.com/Pallinder/go-randomdata"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type RatingController struct {
	Controller
}

//GetRating returns a player's session rating by their Ids
func (*RatingController) GetRating(w http.ResponseWriter, req *http.Request) {

	//Get session from URL pathand player from URL path
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
func (*RatingController) GetRatings(w http.ResponseWriter, req *http.Request) {
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

//*****POST Handlers*****//

func (*RatingController) CreateRating(w http.ResponseWriter, req *http.Request) {

	//Get SessionId from parameter string
	params := mux.Vars(req)
	sessionId := params["SessionId"]
	playerId := req.URL.Query().Get("PlayerId")
	rating := req.URL.Query().Get("Rating")
	comment := req.URL.Query().Get("Comment")
	var responseData response.Response

	//Check and prevent player from submitting another rating for the session if one exists, otherwise insert new rating
	currentRating := db.SelectRating(sessionId, playerId)
	if !currentRating.IsEmpty() {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreateRating",
			Message: "Player has already submitted a rating for the session. Cannot submit more than one rating for a session. Session: " + sessionId + " Player: " + playerId + " rating: " + strconv.Itoa(currentRating.Rating) + " Comment: " + currentRating.Comment,
		})
		return
	}

	ratingInt, err := strconv.Atoi(rating)
	if err != nil {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreateRating",
			Message: "Invalid rating {" + rating + "}. Log: " + err.Error(),
		})
		return
	}

	if ratingInt < 1 || ratingInt > 5 {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreateRating",
			Message: "Rating submitted is not valid. Ratings must be between 1 and 5.",
		})
		return
	}

	timeSubmitted := time.Now()

	//Insert new Player into database
	ok, err := db.InsertNewRating(sessionId, playerId, ratingInt, comment, timeSubmitted)

	if err != nil {
		panic(err)
	}

	//Check if insert was successful and send response
	if ok == true {
		responseData = response.Response{
			Code: http.StatusOK,
			Data: "Rating Successfully submitted for Session ID: " + sessionId + " rating: " + rating + comment,
		}
	} else {
		responseData = response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreateRating",
			Message: "Rating was unable to be submitted for Session ID: " + sessionId + " rating: " + rating + comment,
		}
	}
	response.Write(w, responseData)
}

//GetRandomRating Simulates returning a rating. Nothing is retrieved from or committed to database.
func (*RatingController) GetRandomRating(w http.ResponseWriter, req *http.Request) {

	//Generate random session rating data
	sessionId := xid.New().String()
	playerId := xid.New().String()
	rand.Seed(time.Now().UnixNano())
	rating := 1 + rand.Intn(5-1+1)
	ratingComment := randomdata.Paragraph()

	//Trim random comment data to be less than 512 as that is the limit of the db comment field
	if len(ratingComment) > 511 {
		ratingComment = ratingComment[0:511]
	}

	timeSubmitted := time.Now()
	randomRatingData := models.Rating{PlayerId: playerId, SessionId: sessionId, Rating: rating, Comment: ratingComment, TimeSubmitted: timeSubmitted}
	response.Write(w, response.Response{
		Code: http.StatusOK,
		Data: randomRatingData,
	})
}
