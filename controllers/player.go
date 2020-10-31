package controllers

import (
	"ManOnTheMoonReviewService/controllers/response"
	"ManOnTheMoonReviewService/db"
	"ManOnTheMoonReviewService/models"
	"ManOnTheMoonReviewService/util"
	"encoding/json"
	"github.com/Pallinder/go-randomdata"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type PlayerController struct {
	Controller
}

//GetPlayer returns data of a single player by Id
func (p *PlayerController) GetPlayer(w http.ResponseWriter, req *http.Request) {

	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields() // catch unwanted fields
	var player models.Player

	err := d.Decode(&player)
	if err != nil {
		// bad JSON or unrecognized json field
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Get playerId from URL path
	params := mux.Vars(req)
	playerId := params["PlayerId"]

	//Retrieve player data by id
	playerData := db.SelectPlayer(playerId)

	//Check if player was retrieved and send response
	if playerData.PlayerId == "" {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetPlayer",
			Message: "Could not find player using PlayerId: " + playerId,
		})
	} else {
		response.Write(w, response.Response{
			Code: http.StatusOK,
			Data: playerData,
		})
	}
}

func (p *PlayerController) CreatePlayer(w http.ResponseWriter, req *http.Request) {

	//Generate new PlayerId
	playerID := util.NewUUID()

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
	var responseData response.Response
	if ok == true {
		responseData = response.Response{
			Code: http.StatusOK,
			Data: "New Player Successfully created. ID: " + playerID,
		}
	} else {
		responseData = response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreatePlayer",
			Message: "New Player was not created. ID: " + playerID,
		}
	}

	//Send response as JSON
	response.Write(w, responseData)
}

//GetAllPlayers returns all players from the players table
func (p *PlayerController) GetAllPlayers(w http.ResponseWriter, req *http.Request) {

	//Retrieve all players from players table
	playersData := db.SelectAllPlayers()
	var responseData response.Response

	//Check if players were found then return response
	if len(playersData) == 0 {
		responseData = response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetAllPlayers",
			Message: "No players could be found.",
		}
	} else {
		//Return players data as a JSON response
		responseData = response.Response{
			Code: http.StatusOK,
			Data: playersData,
		}
	}

	response.Write(w, responseData)
}
