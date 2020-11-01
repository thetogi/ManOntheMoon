package models

import (
	"ManOnTheMoonReviewService/db"
	"database/sql"
	"fmt"
	"log"
	"time"
)

//Session Rating Properties
type Rating struct {
	Model
	SessionId     string
	PlayerId      string
	Rating        int
	Comment       string
	TimeSubmitted time.Time
}

type Ratings struct {
	Model
	Data    []Rating
	Options Options
}

type Options struct {
	Rating        int
	FilterOperand string
	Recent        bool
}

func (r *Rating) IsEmpty() bool {
	if *r == (Rating{}) {
		return true
	}
	return false
}

func (r *Rating) Retrieve(rating *Rating) error {
	//Retrieve player data by id
	*rating = SelectRating(rating.SessionId, rating.PlayerId)
	return nil
}

func (r *Rating) Create(rating *Rating) (bool, error) {

	//Insert new Player into database
	ok, err := InsertNewRating(
		rating.SessionId,
		rating.PlayerId,
		rating.Rating,
		rating.Comment,
		rating.TimeSubmitted,
	)

	//If there is an error inserting, handle it
	if err != nil {
		panic(err)
	}
	return ok, nil
}

func (r *Rating) RetrieveAll(ratings *Ratings) error {
	//Retrieve player data by id
	ratings.Data = SelectAllRatings(ratings.Options)
	return nil
}

func SelectRating(sessionId string, playerId string) Rating {
	fmt.Println("Executing SELECT: SelectRating")

	sqlStatement := "SELECT r.SessionId, r.PlayerId, r.Rating, r.Comment, r.TimeSubmitted FROM ratings r WHERE r.SessionId = ? AND r.PlayerId = ?"

	stmt, err := db.Db.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	var ratingData Rating
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

func SelectAllRatings(ops Options) []Rating {
	fmt.Println("Executing SELECT: SelectAllRatings")
	var ratings []Rating

	var limitPart string
	var sqlStatement string
	var filterPart string

	var rows *sql.Rows
	var err error

	//Check for returning only recent reviews and build Select clause
	if ops.Recent {
		limitPart = "Limit 20"
	}

	//Check for rating filter and build the filter clause
	if ops.FilterOperand != "" {
		switch ops.FilterOperand {
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
		rows, err = db.Db.Query(sqlStatement, ops.Rating)
	} else if ops.Rating != 0 {
		//Combine parts to build SQL statement
		sqlStatement = "SELECT r.SessionId, r.PlayerId, r.Rating, r.Comment, r.TimeSubmitted FROM ratings r WHERE r.Rating = ? ORDER BY r.TimeSubmitted DESC " + limitPart
		rows, err = db.Db.Query(sqlStatement, ops.Rating)
	} else {
		sqlStatement = "SELECT r.SessionId, r.PlayerId, r.Rating, r.Comment, r.TimeSubmitted FROM ratings r ORDER BY r.TimeSubmitted DESC " + limitPart
		rows, err = db.Db.Query(sqlStatement)
	}

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var rating Rating
	for rows.Next() {
		switch err := rows.Scan(&rating.SessionId, &rating.PlayerId, &rating.Rating, &rating.Comment, &rating.TimeSubmitted); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			ratings = append(ratings, Rating{SessionId: rating.SessionId, PlayerId: rating.PlayerId, Rating: rating.Rating, Comment: rating.Comment, TimeSubmitted: rating.TimeSubmitted})
			fmt.Println(rating.SessionId, rating.PlayerId, rating.Rating, rating.Comment, rating.TimeSubmitted)
		default:
			panic(err)
		}
		// get any error encountered during iteration
		err = rows.Err()
	}
	return ratings
}

func InsertNewRating(sessionId string, playerId string, rating int, comment string, timeSubmitted time.Time) (bool, error) {
	fmt.Println("Executing INSERT: InsertNewRating")

	sqlStatement := "INSERT INTO ratings (`SessionId`,`PlayerId`,`Rating`,`Comment`, `TimeSubmitted`) VALUES ( ?,?,?,?,?)"

	stmt, err := db.Db.Prepare(sqlStatement)
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
