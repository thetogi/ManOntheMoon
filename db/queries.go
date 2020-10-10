package db

import (
	"fmt"

	"database/sql"
	"log"
	"time"
)

//*****Select Queries*******//

//Retrieve game information for a single game
func SelectGameById(gameId string) Game {
	fmt.Println("Executing SELECT: SelectGameById")

	sqlStatement := "SELECT g.GameId, g.Name, g.Version FROM games g WHERE g.GameID = ?"

	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	var Game Game
	switch err := stmt.QueryRow(gameId).Scan(&Game.GameId, &Game.Name, &Game.Version); err {
	case sql.ErrNoRows:
		fmt.Println("No game was found!, GameId: " + gameId)
	case nil:
		fmt.Println(Game.GameId, Game.Name, Game.Version)
	default:
		panic(err)
	}

	defer stmt.Close()

	return Game
}

func SelectPlayerByPlayerId(playerId string) Player {
	fmt.Println("Executing SELECT: GetPlayerByPlayerId")

	sqlStatement := "SELECT p.PlayerId, p.Name, p.TimeRegistered FROM players p WHERE p.PlayerID = ?"

	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	var PlayerData Player
	switch err := stmt.QueryRow(playerId).Scan(&PlayerData.PlayerId, &PlayerData.Name, &PlayerData.TimeRegistered); err {
	case sql.ErrNoRows:
		fmt.Println("No player was found!, PlayerId: " + playerId)
	case nil:
		fmt.Println(PlayerData.PlayerId, PlayerData.Name, PlayerData.TimeRegistered)
	default:
		panic(err)
	}

	defer stmt.Close()

	return PlayerData
}

func SelectAllPlayers() []Player {
	fmt.Println("Executing SELECT: SelectAllPlayers")

	sqlStatement := "SELECT p.PlayerId, p.Name, p.TimeRegistered FROM players p"

	rows, err := Db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var Players []Player
	var PlayerData Player
	for rows.Next() {
		switch err := rows.Scan(&PlayerData.PlayerId, &PlayerData.Name, &PlayerData.TimeRegistered); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			Players = append(Players, Player{PlayerData.PlayerId, PlayerData.Name, PlayerData.TimeRegistered})
			fmt.Println(PlayerData.PlayerId, PlayerData.Name, PlayerData.TimeRegistered)
		default:
			panic(err)
		}
		// get any error encountered during iteration
		err = rows.Err()
	}
	return Players
}

func SelectSessionbyId(sessionId string) Session {

	fmt.Println("Executing SELECT: SelectSessionById")

	sqlStatement := "SELECT s.SessionId, s.PlayerId, s.GameId, s.TimeSessionEnd FROM Sessions s WHERE s.SessionId = ?"

	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	var Session Session
	switch err := stmt.QueryRow(sessionId).Scan(&Session.SessionId, &Session.PlayerId, &Session.GameId, &Session.TimeSessionEnd); err {
	case sql.ErrNoRows:
		fmt.Println("No session was found!, SessionId: " + sessionId)
	case nil:
		fmt.Println(Session.SessionId, Session.PlayerId, Session.GameId, Session.TimeSessionEnd)
	default:
		panic(err)
	}

	defer stmt.Close()

	return Session
}

func SelectAllSessions() []Session {
	fmt.Println("Executing SELECT: SelectAllSessions")
	sqlStatement := "SELECT s.SessionId, s.PlayerId, s.GameId, s.TimeSessionEnd FROM Sessions s"

	rows, err := Db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var Sessions []Session
	var SingleSession Session
	for rows.Next() {
		switch err := rows.Scan(&SingleSession.SessionId, &SingleSession.PlayerId, &SingleSession.GameId, &SingleSession.TimeSessionEnd); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			Sessions = append(Sessions, Session{SingleSession.SessionId, SingleSession.PlayerId, SingleSession.GameId, SingleSession.TimeSessionEnd})
			fmt.Println(SingleSession.SessionId, SingleSession.PlayerId, SingleSession.GameId, SingleSession.TimeSessionEnd)
		default:
			panic(err)
		}
		// get any error encountered during iteration
		err = rows.Err()
	}
	return Sessions
}

func SelectSessionRatingBySessionId(sessionId string, playerId string) SessionRating {
	fmt.Println("Executing SELECT: SelectSessionRatingBySessionId")

	sqlStatement := "SELECT sr.SessionId, sr.Rating, sr.Comment, sr.TimeSubmitted FROM SessionRatings sr WHERE sr.SessionId = ?"

	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	var sessionRatingData SessionRating
	switch err := stmt.QueryRow(sessionId, playerId).Scan(&sessionRatingData.SessionId, &sessionRatingData.Rating, &sessionRatingData.Comment, &sessionRatingData.TimeSubmitted); err {
	case sql.ErrNoRows:
		fmt.Println("No session rating was found!, SessionId: " + sessionId + " PlayerId: " + playerId)
	case nil:
		fmt.Println(sessionRatingData.SessionId, sessionRatingData.Rating, sessionRatingData.Comment, sessionRatingData.TimeSubmitted)
	default:
		panic(err)
	}

	defer stmt.Close()

	return sessionRatingData
}

func SelectAllSessionRatings(rating int, ratingFilter string) []SessionRating {
	fmt.Println("Executing SELECT: SelectAllSessionRatings")
	var SessionRatings []SessionRating

	var sqlStatement string
	var rows *sql.Rows
	var err error
	if ratingFilter != "" {
		sqlStatement = "SELECT sr.SessionId, sr.PlayerId, sr.Rating, sr.Comment, sr.TimeSubmitted FROM SessionRatings sr WHERE sr.Rating = ?"
		rows, err = Db.Query(sqlStatement, rating)
	} else {
		sqlStatement = "SELECT sr.SessionId, sr.PlayerId, sr.Rating, sr.Comment, sr.TimeSubmitted FROM SessionRatings sr"
		rows, err = Db.Query(sqlStatement)
	}

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var Rating SessionRating
	for rows.Next() {
		switch err := rows.Scan(&Rating.SessionId, &Rating.PlayerId, &Rating.Rating, &Rating.Comment, &Rating.TimeSubmitted); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			SessionRatings = append(SessionRatings, SessionRating{Rating.SessionId, Rating.PlayerId, Rating.Rating, Rating.Comment, Rating.TimeSubmitted})
			fmt.Println(Rating.SessionId, Rating.PlayerId, Rating.Rating, Rating.Comment, Rating.TimeSubmitted)
		default:
			panic(err)
		}
		// get any error encountered during iteration
		err = rows.Err()
	}
	return SessionRatings
}

//*****Insert Queries*******//

func InsertNewPlayer(playerId string, playerName string, TimeRegistered time.Time) (bool, error) {
	fmt.Println("Executing INSERT: InsertNewPlayer")
	sqlStatement := "INSERT INTO players (`PlayerId`,`Name`,`TimeRegistered`) VALUES (?,?,?)"
	insert, err := Db.Query(sqlStatement, playerId, playerName, TimeRegistered)

	// if there is an error inserting, handle it
	if err != nil {
		return false, err
	}
	defer insert.Close()

	return true, err
}

func InsertNewSession(sessionId string, gameId string, playerId string, timeSessionEnd time.Time) (bool, error) {
	fmt.Println("Executing INSERT: InsertNewSession")
	sqlStatement := "INSERT INTO Sessions (`SessionId`,`GameId`,`PlayerId`,`TimeSessionEnd`) VALUES (?,?,?,?)"
	insert, err := Db.Query(sqlStatement, sessionId, gameId, playerId, timeSessionEnd)

	// if there is an error inserting, handle it
	if err != nil {
		return false, err
	}
	defer insert.Close()

	return true, err

}

func InsertNewSessionRating(sessionId string, playerId string, rating int, comment string, timeSubmitted time.Time) (bool, error) {
	fmt.Println("Executing INSERT: InsertNewSessionRating")
	sqlStatement := "INSERT INTO SessionRatings (`SessionId`,`PlayerId`,`Rating`,`Comment`, `TimeSubmitted`) VALUES ( ?,?,?,?,?)"
	insert, err := Db.Query(sqlStatement, sessionId, playerId, rating, comment, timeSubmitted)

	// if there is an error inserting, handle it
	if err != nil {
		return false, err
	}
	defer insert.Close()

	return true, err
}
