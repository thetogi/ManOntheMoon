package main

import (
	"ManOnTheMoonReviewService/db"
	"ManOnTheMoonReviewService/models"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/rs/xid"
	"math/rand"
	"time"
)

//Generates random data for the database
func main() {

	//Create 10 random Players, Sessions, and Ratings

	//Generate random session rating data
	playerCount := 10
	sessionCount := 1 + rand.Intn(15-1+1)

	for i := playerCount; i > 0; i-- {

		playerData := models.Player{PlayerId: xid.New().String(), Name: randomdata.FullName(randomdata.RandomGender), TimeRegistered: time.Now()}

		db.InsertNewPlayer(playerData.PlayerId, playerData.Name, playerData.TimeRegistered)

		for i := sessionCount; i > 0; i-- {
			sessionData := models.Session{SessionId: xid.New().String(), PlayerId: playerData.PlayerId, TimeSessionEnd: time.Now()}

			db.InsertNewSession(sessionData.SessionId, sessionData.PlayerId, sessionData.TimeSessionEnd)

			fmt.Println("Created session: ", sessionData.SessionId, " for: ", playerData.Name, " Player Id: ", playerData.PlayerId)

			ratingData := models.Rating{
				SessionId:     sessionData.SessionId,
				PlayerId:      playerData.PlayerId,
				Rating:        1 + rand.Intn(5-1+1),
				Comment:       randomdata.Paragraph(),
				TimeSubmitted: time.Now()}

			if len(ratingData.Comment) > 511 {
				ratingData.Comment = ratingData.Comment[0:511]
			}

			db.InsertNewSessionRating(ratingData.SessionId, ratingData.PlayerId, ratingData.Rating, ratingData.Comment, ratingData.TimeSubmitted)
			fmt.Println("Created session rating for: ", playerData.Name, " Player Id: ", playerData.PlayerId, " session Id: ", sessionData.SessionId, " rating: ", ratingData.Rating)
		}
		fmt.Println("Created player: ", playerData.Name, " Player Id: ", playerData.PlayerId)
	}

	playerCount2 := 10
	sessionCount2 := 1 + rand.Intn(15-1+1)

	for i := playerCount2; i > 0; i-- {

		playerData := models.Player{PlayerId: xid.New().String(), Name: randomdata.FullName(randomdata.RandomGender), TimeRegistered: time.Now()}

		db.InsertNewPlayer(playerData.PlayerId, playerData.Name, playerData.TimeRegistered)

		for i := sessionCount2; i > 0; i-- {
			sessionData := models.Session{SessionId: xid.New().String(), PlayerId: playerData.PlayerId, TimeSessionEnd: time.Now()}

			db.InsertNewSession(sessionData.SessionId, sessionData.PlayerId, sessionData.TimeSessionEnd)

			fmt.Println("Created session no rating: ", sessionData.SessionId, " for: ", playerData.Name, " Player Id: ", playerData.PlayerId)
		}
		fmt.Println("Created player no rating: ", playerData.Name, " Player Id: ", playerData.PlayerId)
	}
}
