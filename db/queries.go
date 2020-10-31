package db

//This file contains all the database queries that are used in the service
//the database/sql library and driver will sanitize inputs using parameterized queries to prevent SQL injection.

import (
	"ManOnTheMoonReviewService/models"
	"fmt"

	"database/sql"
	"log"
	"time"
)

//*****Select Queries*******//

func SelectPlayer(playerId string) models.Player {
	fmt.Println("Executing SELECT: GetPlayerByPlayerId")

	sqlStatement := "SELECT p.PlayerId, p.Name, p.TimeRegistered FROM players p WHERE p.PlayerID = ?"

	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	var PlayerData models.Player
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

func SelectAllPlayers() []models.Player {
	fmt.Println("Executing SELECT: SelectAllPlayers")

	sqlStatement := "SELECT p.PlayerId, p.Name, p.TimeRegistered FROM players p"

	rows, err := Db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var Players []models.Player
	var PlayerData models.Player
	for rows.Next() {
		switch err := rows.Scan(&PlayerData.PlayerId, &PlayerData.Name, &PlayerData.TimeRegistered); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			Players = append(Players, models.Player{PlayerId: PlayerData.PlayerId, Name: PlayerData.Name, TimeRegistered: PlayerData.TimeRegistered})
			fmt.Println(PlayerData.PlayerId, PlayerData.Name, PlayerData.TimeRegistered)
		default:
			panic(err)
		}
		// get any error encountered during iteration
		err = rows.Err()
	}
	return Players
}

func SelectSession(sessionId string) models.Session {

	fmt.Println("Executing SELECT: SelectSessionById")

	sqlStatement := "SELECT s.SessionId, s.PlayerId, s.TimeSessionEnd FROM Sessions s WHERE s.SessionId = ?"

	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	var Session models.Session
	switch err := stmt.QueryRow(sessionId).Scan(&Session.SessionId, &Session.PlayerId, &Session.TimeSessionEnd); err {
	case sql.ErrNoRows:
		fmt.Println("No session was found!, SessionId: " + sessionId)
	case nil:
		fmt.Println(Session.SessionId, Session.PlayerId, Session.TimeSessionEnd)
	default:
		panic(err)
	}

	defer stmt.Close()

	return Session
}

func SelectAllSessions() []models.Session {
	fmt.Println("Executing SELECT: SelectAllSessions")
	sqlStatement := "SELECT s.SessionId, s.PlayerId, s.TimeSessionEnd FROM Sessions s"

	rows, err := Db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var Sessions []models.Session
	var SingleSession models.Session
	for rows.Next() {
		switch err := rows.Scan(&SingleSession.SessionId, &SingleSession.PlayerId, &SingleSession.TimeSessionEnd); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			Sessions = append(Sessions, models.Session{SingleSession.SessionId, SingleSession.PlayerId, SingleSession.TimeSessionEnd})
			fmt.Println(SingleSession.SessionId, SingleSession.PlayerId, SingleSession.TimeSessionEnd)
		default:
			panic(err)
		}
		// get any error encountered during iteration
		err = rows.Err()
	}
	return Sessions
}

func SelectRating(sessionId string, playerId string) models.Rating {
	fmt.Println("Executing SELECT: SelectRating")

	sqlStatement := "SELECT r.SessionId, r.PlayerId, r.Rating, r.Comment, r.TimeSubmitted FROM ratings r WHERE r.SessionId = ? AND r.PlayerId = ?"

	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	var ratingData models.Rating
	switch err := stmt.QueryRow(sessionId, playerId).Scan(&ratingData.SessionId, &ratingData.PlayerId, &ratingData.Rating, &ratingData.Comment, &ratingData.TimeSubmitted); err {
	case sql.ErrNoRows:
		fmt.Println("No session rating was found!, SessionId: " + sessionId + " PlayerId: " + playerId)
	case nil:
		fmt.Println(ratingData.SessionId, ratingData.Rating, ratingData.Comment, ratingData.TimeSubmitted)
	default:
		panic(err)
	}

	defer stmt.Close()

	return ratingData
}

func SelectAllRatings(rating int, ratingFilterOp string, recentFlag bool) []models.Rating {
	fmt.Println("Executing SELECT: SelectAllRatings")
	var ratings []models.Rating

	var limitPart string
	var sqlStatement string
	var filterPart string

	var rows *sql.Rows
	var err error

	//Check for returning only recent reviews and build Select clause
	if recentFlag {
		limitPart = "Limit 20"
	}

	//Check for rating filter and build the filter clause
	if ratingFilterOp != "" {
		switch ratingFilterOp {
		case ">":
			filterPart = "WHERE r.Rating > ?"
		case ">=":
			filterPart = "WHERE r.Rating >= ?"
		case "<":
			filterPart = "WHERE r.Rating < ?"
		case "<=":
			filterPart = "WHERE r.Rating <= ?"
		}
		//Combine parts to build SQL statement
		sqlStatement = "SELECT r.SessionId, r.PlayerId, r.Rating, r.Comment, r.TimeSubmitted FROM ratings r " + filterPart + " ORDER BY r.TimeSubmitted DESC " + limitPart
		rows, err = Db.Query(sqlStatement, rating)
	} else if rating != 0 {
		//Combine parts to build SQL statement
		sqlStatement = "SELECT r.SessionId, r.PlayerId, r.Rating, r.Comment, r.TimeSubmitted FROM ratings r WHERE r.Rating = ? ORDER BY r.TimeSubmitted DESC " + limitPart
		rows, err = Db.Query(sqlStatement, rating)
	} else {
		sqlStatement = "SELECT r.SessionId, r.PlayerId, r.Rating, r.Comment, r.TimeSubmitted FROM ratings r ORDER BY r.TimeSubmitted DESC " + limitPart
		rows, err = Db.Query(sqlStatement)
	}

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var Rating models.Rating
	for rows.Next() {
		switch err := rows.Scan(&Rating.SessionId, &Rating.PlayerId, &Rating.Rating, &Rating.Comment, &Rating.TimeSubmitted); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			ratings = append(ratings, models.Rating{SessionId: Rating.SessionId, PlayerId: Rating.PlayerId, Rating: Rating.Rating, Comment: Rating.Comment, TimeSubmitted: Rating.TimeSubmitted})
			fmt.Println(Rating.SessionId, Rating.PlayerId, Rating.Rating, Rating.Comment, Rating.TimeSubmitted)
		default:
			panic(err)
		}
		// get any error encountered during iteration
		err = rows.Err()
	}
	return ratings
}

//*****Insert Queries*******//

func InsertNewPlayer(playerId string, playerName string, TimeRegistered time.Time) (bool, error) {
	fmt.Println("Executing INSERT: InsertNewPlayer")

	sqlStatement := "INSERT INTO players (`PlayerId`,`Name`,`TimeRegistered`) VALUES (?,?,?)"

	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(playerId, playerName, TimeRegistered)

	// if there is an error inserting, handle it
	if err != nil {
		return false, err
	}

	return true, err
}

func InsertNewSession(sessionId string, playerId string, timeSessionEnd time.Time) (bool, error) {
	fmt.Println("Executing INSERT: InsertNewSession")

	sqlStatement := "INSERT INTO Sessions (`SessionId`,`PlayerId`,`TimeSessionEnd`) VALUES (?,?,?)"

	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sessionId, playerId, timeSessionEnd)

	// if there is an error inserting, handle it
	if err != nil {
		return false, err
	}

	return true, err
}

func InsertNewRating(sessionId string, playerId string, rating int, comment string, timeSubmitted time.Time) (bool, error) {
	fmt.Println("Executing INSERT: InsertNewRating")

	sqlStatement := "INSERT INTO ratings (`SessionId`,`PlayerId`,`Rating`,`Comment`, `TimeSubmitted`) VALUES ( ?,?,?,?,?)"

	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sessionId, playerId, rating, comment, timeSubmitted)

	// if there is an error inserting, handle it
	if err != nil {
		return false, err
	}

	return true, err
}
