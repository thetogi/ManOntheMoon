package seed

import (
	"ManOnTheMoonReviewService/models"
	"ManOnTheMoonReviewService/util"
	"github.com/Pallinder/go-randomdata"
	"time"
)

type Seeder struct {
	Player  models.Player
	Session models.Session
	Rating  models.Rating
}

func (*Seeder) Generate() {

}

func (*Seeder) NewPlayer() models.Player {

	return models.Player{PlayerId: util.NewUUID(), Name: randomdata.FullName(randomdata.RandomGender), TimeRegistered: time.Now()}

}

func (*Seeder) NewSession(playerId string) models.Session {
	return models.Session{
		SessionId: util.NewUUID(),
		PlayerId:  playerId,
	}
}

func (*Seeder) NewRating(playerId string, sessionId string) models.Rating {
	return models.Rating{
		SessionId:     "",
		PlayerId:      "",
		Rating:        0,
		Comment:       "",
		TimeSubmitted: time.Time{},
	}
}
