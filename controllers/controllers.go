package controllers

import (
	"ManOnTheMoonReviewService/db"
	"ManOnTheMoonReviewService/util"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//Game Structs
type Game struct {
	GameId  string
	Name    string
	Version float32
}

//Player Structs
type Player struct {
	PlayerId       string
	Name           string
	TimeRegistered time.Time
}

type Players struct {
	Data []Player
}

//Session Structs
type Session struct {
	SessionId      string
	PlayerId       string
	GameId         string
	TimeSessionEnd time.Time
}

type Sessions struct {
	Data []Session
}

type SessionRatingData struct {
	PlayerId      string
	Session       string
	Rating        int
	Comment       string
	DateSubmitted time.Time
}

//Session Rating Structs
type SessionRating struct {
	SessionId     string
	Rating        int
	Comment       string
	TimeSubmitted time.Time
}

type SessionRatings struct {
	Data []SessionRating
}

//Http Response Structs
type ResponsePostSessionRating struct {
	Status  string
	Message string
}

type ResponsePost struct {
	Status  string
	Message string
}

type ResponseGet struct {
	Status  string
	Message string
}

func ShowEnvToggle(w http.ResponseWriter, req *http.Request) {

	files := []string{
		"./ui/html/envToggle.layout.tmpl",
		"./ui/html/envToggle.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "Internal Server Error", 500)

	}
	if util.EnvShowStatus() == false {
		util.ToggleEnvShow()
	}

	err = ts.Execute(w, util.EnvData())
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}

func Home(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		w.WriteHeader(404)
		w.Write([]byte("Not Found !!!"))
		return
	}
	newSessionId := xid.New()
	//Store session ID in DB

	// Redirect to new session
	http.Redirect(w, req, "http://localhost:8080/GameSession/"+newSessionId.String(), http.StatusSeeOther)

	return
	//files := []string{
	//	"./ui/html/home.page.tmpl",
	//	"./ui/html/base.layout.tmpl",
	//	"./ui/html/footer.partial.tmpl",
	//}
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	log.Printf(err.Error())
	//	http.Error(w, "Internal Server Error", 500)
	//
	//}
	//
	//err = ts.Execute(w, util.EnvData())
	//if err != nil {
	//	log.Printf(err.Error())
	//	http.Error(w, "Internal Server Error", 500)
	//}
}

func PostNewReview(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	//playerID, _ := strconv.Atoi(req.PostFormValue("playerID"))
	//session, _ := strconv.Atoi(req.PostFormValue("session"))
	//SessionRating, _ := strconv.Atoi(req.PostFormValue("rating"))
	//SessionRatingComment := req.PostFormValue("comment")

	// Commit to database

	//Return message
	//data := SessionRatingData{PlayerId:playerID, Session: session, Rating: SessionRating, Comment: SessionRatingComment}
	responseData := ResponsePostSessionRating{Status: "200", Message: "Success"}
	sessionRatingJson, err := json.Marshal(responseData)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(sessionRatingJson)
}

func GetRatingRandom(w http.ResponseWriter, req *http.Request) {

	playerID := xid.New().String()
	session := xid.New().String()
	rand.Seed(time.Now().UnixNano())
	SessionRating := 1 + rand.Intn(5-1+1)
	SessionRatingComment := "hello world"
	DateSubmitted := time.Now()
	data := SessionRatingData{PlayerId: playerID, Session: session, Rating: SessionRating, Comment: SessionRatingComment, DateSubmitted: DateSubmitted}

	sessionRatingJson, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(sessionRatingJson)
}

func PostGameCreateSql(w http.ResponseWriter, req *http.Request) {

	gameId := xid.New().String()
	gameName := "ManOnTheMoon"
	gameVersion := "1.0"
	sqlStatement := "INSERT INTO games (`GameId`,`Name`,`Version`) VALUES ( ?, ?, ?)"
	insert, err := db.Db.Query(sqlStatement, gameId, gameName, gameVersion)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	data := ResponsePost{Status: "OK", Message: "New GAme Successfully created. ID: " + gameId}

	ResponsePost, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ResponsePost)

}

func GetGameByIdSql(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	gameId := params["GameId"]

	sqlStatement := "SELECT g.GameId, g.Name, g.Version FROM games g WHERE g.GameID = ?"
	var Game Game
	row := db.Db.QueryRow(sqlStatement, gameId)
	switch err := row.Scan(&Game.GameId, &Game.Name, &Game.Version); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(Game.GameId, Game.Name, Game.Version)
	default:
		panic(err)
	}

	ResponsePost, err := json.Marshal(Game)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ResponsePost)
}

func PostPlayerCreateSql(w http.ResponseWriter, req *http.Request) {

	playerID := xid.New().String()
	playerName := randomdata.FullName(randomdata.RandomGender)
	PlayerTimeRegistered := time.Now()
	sqlStatement := "INSERT INTO players (`PlayerId`,`Name`,`TimeRegistered`) VALUES ( ?, ?, ?)"
	insert, err := db.Db.Query(sqlStatement, playerID, playerName, PlayerTimeRegistered)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	data := ResponsePost{Status: "OK", Message: "New Player Successfully created. ID: " + playerID}

	ResponsePost, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ResponsePost)

}

func GetPlayerByIdSql(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	playerId := params["PlayerId"]

	sqlStatement := "SELECT p.PlayerId, p.Name, p.TimeRegistered FROM players p WHERE p.PlayerID = ?"
	var PlayerData Player
	row := db.Db.QueryRow(sqlStatement, playerId)
	switch err := row.Scan(&PlayerData.PlayerId, &PlayerData.Name, &PlayerData.TimeRegistered); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(PlayerData.PlayerId, PlayerData.Name, PlayerData.TimeRegistered)
	default:
		panic(err)
	}

	ResponsePost, err := json.Marshal(PlayerData)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ResponsePost)
}

func GetAllPlayersSql(w http.ResponseWriter, req *http.Request) {
	var Players Players
	sqlStatement := "SELECT p.PlayerId, p.Name, p.TimeRegistered FROM players p"
	var PlayerData Player
	rows, err := db.Db.Query(sqlStatement)

	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		switch err := rows.Scan(&PlayerData.PlayerId, &PlayerData.Name, &PlayerData.TimeRegistered); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			Players.Data = append(Players.Data, Player{PlayerData.PlayerId, PlayerData.Name, PlayerData.TimeRegistered})
			fmt.Println(PlayerData.PlayerId, PlayerData.Name, PlayerData.TimeRegistered)
		default:
			panic(err)
		}
		// get any error encountered during iteration
		err = rows.Err()
	}
	ResponseGet, err := json.Marshal(Players.Data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ResponseGet)
}

func PostSessionCreateSql(w http.ResponseWriter, req *http.Request) {
	sessionId := xid.New().String()
	playerID := req.URL.Query().Get("PlayerId")
	gameId := req.URL.Query().Get("GameId")
	timeSessionEnd, _ := time.Parse(time.RFC822, req.URL.Query().Get("TimeSessionEnd"))
	sqlStatement := "INSERT INTO Sessions (`SessionId`,`GameId`,`PlayerId`,`TimeSessionEnd`) VALUES ( ?, ?, ?,?)"
	insert, err := db.Db.Query(sqlStatement, sessionId, gameId, playerID, timeSessionEnd)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	data := ResponsePost{Status: "OK", Message: "New Session Successfully created. ID: " + sessionId}

	ResponsePost, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ResponsePost)

}

func GetSessionByIdSql(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	playerId := params["SessionId"]

	sqlStatement := "SELECT s.SessionId, s.PlayerId, s.GameId, s.TimeSessionEnd FROM Sessions s WHERE s.SessionId = ?"
	var Session Session
	row := db.Db.QueryRow(sqlStatement, playerId)
	switch err := row.Scan(&Session.SessionId, &Session.PlayerId, &Session.GameId, &Session.TimeSessionEnd); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(Session.SessionId, Session.PlayerId, Session.GameId, Session.TimeSessionEnd)
	default:
		panic(err)
	}

	ResponsePost, err := json.Marshal(Session)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ResponsePost)
}

func GetAllSessionsSql(w http.ResponseWriter, req *http.Request) {
	var Sessions Sessions
	sqlStatement := "SELECT s.SessionId, s.PlayerId, s.GameId, s.TimeSessionEnd FROM Sessions s"
	var SingleSession Session
	rows, err := db.Db.Query(sqlStatement)

	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		switch err := rows.Scan(&SingleSession.SessionId, &SingleSession.PlayerId, &SingleSession.GameId, &SingleSession.TimeSessionEnd); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			Sessions.Data = append(Sessions.Data, Session{SingleSession.SessionId, SingleSession.PlayerId, SingleSession.GameId, SingleSession.TimeSessionEnd})
			fmt.Println(SingleSession.SessionId, SingleSession.PlayerId, SingleSession.GameId, SingleSession.TimeSessionEnd)
		default:
			panic(err)
		}
		// get any error encountered during iteration
		err = rows.Err()
	}
	ResponseGet, err := json.Marshal(Sessions.Data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ResponseGet)
}

func PostSessionRatingCreateSql(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	sessionId := params["SessionId"]
	req.ParseForm()
	rating, _ := strconv.Atoi(req.FormValue("Rating"))
	comment := req.Form.Get("Comment")
	timeSubmitted := time.Now()
	sqlStatement := "INSERT INTO SessionRatings (`SessionId`,`Rating`,`Comment`, `TimeSubmitted`) VALUES ( ?, ?, ?,?)"
	insert, err := db.Db.Query(sqlStatement, sessionId, rating, comment, timeSubmitted)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	data := ResponsePost{Status: "OK", Message: "New Rating Successfully submitted for Session ID: " + sessionId + " rating: " + strconv.Itoa(rating) + comment}

	ResponsePost, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ResponsePost)

}

func GetSessionRatingBySessionIdSql(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	sessionId := params["SessionId"]

	sqlStatement := "SELECT sr.SessionId, sr.Rating, sr.Comment, sr.TimeSubmitted FROM SessionRatings sr WHERE sr.SessionId = ?"
	var SessionRatingData SessionRating
	row := db.Db.QueryRow(sqlStatement, sessionId)
	switch err := row.Scan(&SessionRatingData.SessionId, &SessionRatingData.Rating, &SessionRatingData.Comment, &SessionRatingData.TimeSubmitted); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(SessionRatingData.SessionId, SessionRatingData.Rating, SessionRatingData.Comment, SessionRatingData.TimeSubmitted)
	default:
		panic(err)
	}

	ResponsePost, err := json.Marshal(SessionRatingData)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ResponsePost)
}

func GetAllSessionRatingsSql(w http.ResponseWriter, req *http.Request) {

	ratingFilter := req.URL.Query().Get("Rating")

	//ratingFilterOp := req.URL.Query().Get("Op")

	var SessionRatings SessionRatings

	var sqlStatement string

	if ratingFilter != "" {
		sqlStatement = "SELECT sr.SessionId, sr.Rating, sr.Comment, sr.TimeSubmitted FROM SessionRatings sr WHERE sr.Rating = ?"
	} else {
		sqlStatement = "SELECT sr.SessionId, sr.Rating, sr.Comment, sr.TimeSubmitted FROM SessionRatings sr"
	}

	var SingleRating SessionRating
	rows, err := db.Db.Query(sqlStatement)

	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		switch err := rows.Scan(&SingleRating.SessionId, &SingleRating.Rating, &SingleRating.Comment, &SingleRating.TimeSubmitted); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			SessionRatings.Data = append(SessionRatings.Data, SessionRating{SingleRating.SessionId, SingleRating.Rating, SingleRating.Comment, SingleRating.TimeSubmitted})
			fmt.Println(SingleRating.SessionId, SingleRating.Rating, SingleRating.Comment, SingleRating.TimeSubmitted)
		default:
			panic(err)
		}
		// get any error encountered during iteration
		err = rows.Err()
	}
	ResponseGet, err := json.Marshal(SessionRatings.Data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ResponseGet)
}
