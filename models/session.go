package models

import (
	"ManOnTheMoonReviewService/db"
	"ManOnTheMoonReviewService/util"
	"database/sql"
	"fmt"
	"log"
	"time"
)

//Session Properties
type Session struct {
	Model
	SessionId      string
	PlayerId       string
	TimeSessionEnd time.Time
}

type Sessions []Session

func (s *Session) Retrieve(session *Session) error {
	//Retrieve player data by id
	*session = SelectSession(session.SessionId)
	return nil
}

func (s *Session) Create(session *Session) (bool, error) {

	session.SessionId = util.NewUUID()
	session.TimeSessionEnd = time.Now()

	//Insert new Player into database
	ok, err := InsertNewSession(session.SessionId, session.PlayerId, session.TimeSessionEnd)

	//If there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	return ok, nil
}

func (s *Session) RetrieveAll(sessions *Sessions) error {
	//Retrieve player data by id
	*sessions = SelectAllSessions()
	return nil
}

func SelectSession(sessionId string) Session {

	fmt.Println("Executing SELECT: SelectSessionById")

	sqlStatement := "SELECT s.SessionId, s.PlayerId, s.TimeSessionEnd FROM Sessions s WHERE s.SessionId = ?"

	stmt, err := db.Db.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	var Session Session
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

func SelectAllSessions() []Session {
	fmt.Println("Executing SELECT: SelectAllSessions")
	sqlStatement := "SELECT s.SessionId, s.PlayerId, s.TimeSessionEnd FROM sessions s"

	rows, err := db.Db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var sessions []Session
	var singleSession Session
	for rows.Next() {
		switch err := rows.Scan(&singleSession.SessionId, &singleSession.PlayerId, &singleSession.TimeSessionEnd); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			sessions = append(sessions, Session{SessionId: singleSession.SessionId, PlayerId: singleSession.PlayerId, TimeSessionEnd: singleSession.TimeSessionEnd})
			fmt.Println(singleSession.SessionId, singleSession.PlayerId, singleSession.TimeSessionEnd)
		default:
			panic(err)
		}
		// get any error encountered during iteration
		err = rows.Err()
	}
	return sessions
}

func InsertNewSession(sessionId string, playerId string, timeSessionEnd time.Time) (bool, error) {
	fmt.Println("Executing INSERT: InsertNewSession")

	sqlStatement := "INSERT INTO Sessions (`SessionId`,`PlayerId`,`TimeSessionEnd`) VALUES (?,?,?)"

	stmt, err := db.Db.Prepare(sqlStatement)
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
