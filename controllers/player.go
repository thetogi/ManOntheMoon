package controllers

import (
	"ManOnTheMoonReviewService/controllers/response"
	"ManOnTheMoonReviewService/models"
	"ManOnTheMoonReviewService/util"
	"net/http"
)

type PlayerController struct {
	Controller
	Player  models.Player
	Players models.Players
}

//GetPlayer returns data of a single player by Id
func (p *PlayerController) GetPlayer(w http.ResponseWriter, req *http.Request) {

	var player models.Player

	err := util.ParseRequestBody(w, req, &player)
	if err != nil {
		return
	}

	if player.PlayerId == "" {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetPlayer",
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

	//Use model to retrieve player
	p.Player.Retrieve(&player)

	//Check if player was retrieved and send response
	if player.IsEmpty() {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetPlayer",
			Message: "Could not find player",
			Errors:  map[string]string{"PlayerId": player.PlayerId},
		})
	} else {
		response.Write(w, response.Response{
			Code: http.StatusOK,
			Data: player,
		})
	}
}

func (p *PlayerController) CreatePlayer(w http.ResponseWriter, req *http.Request) {

	var newPlayer models.Player

	err := util.ParseRequestBody(w, req, &newPlayer)
	if err != nil {
		return
	}

	if newPlayer.Name == "" {
		response.Write(w, response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreatePlayer",
			Message: "Name cannot be blank",
			Errors:  map[string]string{"Player Name": newPlayer.Name},
		})
		return
	}

	ok, err := p.Player.Create(&newPlayer)

	//Check if insert was successful and send response
	var responseData response.Response
	if ok {
		responseData = response.Response{
			Code: http.StatusOK,
			Data: struct {
				Message string
				Data    models.Player
			}{
				"New Player Successfully created. ID: " + newPlayer.PlayerId,
				newPlayer,
			},
		}
	} else {
		responseData = response.Response{
			Code:    http.StatusBadRequest,
			Action:  "CreatePlayer",
			Message: "Failed to create player",
			Errors:  map[string]string{"Player Name": newPlayer.Name},
		}
	}

	//Send response as JSON
	response.Write(w, responseData)
}

//GetAllPlayers returns all players from the players table
func (p *PlayerController) GetAllPlayers(w http.ResponseWriter, req *http.Request) {

	//Retrieve all players from players table
	p.Players.RetrieveAll()
	var responseData response.Response

	//Check if players were found then return response
	if p.Players.Count() == 0 {
		responseData = response.Response{
			Code:    http.StatusBadRequest,
			Action:  "GetAllPlayers",
			Message: "No players could be found.",
			Errors:  nil,
		}
	} else {
		//Return players data as a JSON response
		responseData = response.Response{
			Code: http.StatusOK,
			Data: p.Players,
		}
	}
	response.Write(w, responseData)
}
