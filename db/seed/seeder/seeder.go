package seed

import (
	"ManOnTheMoonReviewService/models"
	"ManOnTheMoonReviewService/util"
	"errors"
	"github.com/Pallinder/go-randomdata"
	"math/rand"
	"time"
)

type Seeder struct {
	Player      models.Player
	Sessions    []models.Session
	Ratings     []models.Rating
	maxSessions int
	maxRatings  int
}

func (s *Seeder) Generate(sessions int, ratings int) (err error) {
	if ratings > sessions {
		return errors.New("cannot generate more ratings than sessions")
	}

	if ratings != 0 && sessions == 0 {
		return errors.New("cannot generate rating if no sessions are generated")
	}

	s.Player = s.NewPlayer()
	for i, j := 0, 0; i < sessions; i, j = i+1, j+1 {
		s.Sessions = append(s.Sessions, s.NewSession(s.Player.PlayerId))

		if j < ratings {
			s.Ratings = append(s.Ratings, s.NewRating(s.Player.PlayerId, s.Sessions[i].SessionId))
		}

	}
	return nil
}

func (s *Seeder) NewPlayer() models.Player {

	return models.Player{
		PlayerId:       util.NewUUID(),
		Name:           randomdata.FullName(randomdata.RandomGender),
		TimeRegistered: time.Now(),
	}

}

func (s *Seeder) NewSession(playerId string) models.Session {
	return models.Session{
		SessionId: util.NewUUID(),
		PlayerId:  playerId,
	}
}

func (*Seeder) NewRating(playerId string, sessionId string) models.Rating {
	return models.Rating{
		SessionId:     sessionId,
		PlayerId:      playerId,
		Rating:        1 + rand.Intn(5-1+1),
		Comment:       generateComment(),
		TimeSubmitted: time.Now(),
	}
}

func MockPlayerData() models.Player {
	return models.Player{
		PlayerId:       util.NewUUID(),
		Name:           randomdata.FullName(randomdata.RandomGender),
		TimeRegistered: time.Now(),
	}
}

func MockSessionData() models.Session {
	return models.Session{
		SessionId: util.NewUUID(),
		PlayerId:  util.NewUUID(),
	}
}

func MockRatingData() models.Rating {
	dan := models.Rating{
		SessionId:     util.NewUUID(),
		PlayerId:      util.NewUUID(),
		Rating:        1 + rand.Intn(5-1+1),
		Comment:       generateComment(),
		TimeSubmitted: time.Now(),
	}

	return dan

}

func generateComment() string {
	comment := randomdata.Paragraph()

	//Keep comments under 512 characters
	if len(comment) > 511 {
		comment = comment[0:511]
	}
	return comment
}
