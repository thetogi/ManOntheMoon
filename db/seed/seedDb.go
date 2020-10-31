package main

import (
	"ManOnTheMoonReviewService/db"
	seed "ManOnTheMoonReviewService/db/seed/seeder"
	"fmt"
	"log"
	"math/rand"
)

//Generates random data for the database
func main() {

	genPlayers := 100

	for p := genPlayers; p > 0; p-- {

		var seeder seed.Seeder
		genSessions := 1 + rand.Intn(15-1+1)
		seeder.Generate(genSessions, genSessions)

		ok, err := db.InsertNewPlayer(seeder.Player.PlayerId, seeder.Player.Name, seeder.Player.TimeRegistered)
		if !ok {
			log.Fatal("Failed to insert new player. Error: " + err.Error())
		}

		fmt.Println("Created player: ", seeder.Player.Name, " Player Id: ", seeder.Player.PlayerId)

		for s := 0; s < len(seeder.Sessions); s++ {
			db.InsertNewSession(seeder.Sessions[s].SessionId, seeder.Sessions[s].PlayerId, seeder.Sessions[s].TimeSessionEnd)
			if !ok {
				log.Fatal("Failed to insert new session. Error: " + err.Error())
			}
			fmt.Println("Created session: ", seeder.Sessions[s].SessionId, " for: ", seeder.Player.Name, " Player Id: ", seeder.Player.PlayerId)
		}

		for r := 0; r < len(seeder.Ratings); r++ {
			db.InsertNewRating(seeder.Ratings[r].SessionId, seeder.Ratings[r].PlayerId, seeder.Ratings[r].Rating, seeder.Ratings[r].Comment, seeder.Ratings[r].TimeSubmitted)
			if !ok {
				log.Fatal("Failed to insert new rating. Error: " + err.Error())
			}
			fmt.Println("Created rating for: ", seeder.Player.Name, " Player Id: ", seeder.Player.PlayerId, " session Id: ", seeder.Ratings[r].SessionId, " rating: ", seeder.Ratings[r].Rating)
		}
	}
}
