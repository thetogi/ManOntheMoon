package models

import (
	"ManOnTheMoonReviewService/db"
	"ManOnTheMoonReviewService/util"
	"database/sql"
	"fmt"
	"log"
	"time"
)

//Player Properties
type Player struct {
	Model
	PlayerId       string    `json:"playerId"`
	Name           string    `json:"name"`
	TimeRegistered time.Time `json:"timeRegistered"`
}

type Players []Player

func (p *Player) Create(player *Player) (bool, error) {

	//Generate new PlayerId
	player.PlayerId = util.NewUUID()

	//Track Player being registered as current date and time
	player.TimeRegistered = time.Now()

	//Insert new Player into database
	ok, err := InsertNewPlayer(
		player.PlayerId,
		player.Name,
		player.TimeRegistered,
	)

	//If there is an error inserting, handle it
	if err != nil {
		panic(err)
	}
	return ok, nil
}

func (p *Player) Retrieve(player *Player) error {
	//Retrieve player data by id
	*player = SelectPlayer(player.PlayerId)
	return nil
}

func (p *Player) IsEmpty() bool {
	if *p == (Player{}) {
		return true
	}
	return false
}

func (p *Player) IsValid() bool {
	if p.Name == "" || !util.IsValidUUID(p.PlayerId) || p.TimeRegistered.IsZero() {
		return false
	}
	return true
}
func (p *Players) RetrieveAll() error {
	//Retrieve all players
	*p = SelectAllPlayers()
	return nil
}

func (p *Players) Count() int {
	return len(*p)
}

func SelectPlayer(playerId string) Player {
	fmt.Println("Executing SELECT: GetPlayerByPlayerId")

	sqlStatement := "SELECT p.PlayerId, p.Name, p.TimeRegistered FROM players p WHERE p.PlayerID = ?"

	stmt, err := db.Db.Prepare(sqlStatement)
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

	rows, err := db.Db.Query(sqlStatement)

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
			Players = append(Players, Player{PlayerId: PlayerData.PlayerId, Name: PlayerData.Name, TimeRegistered: PlayerData.TimeRegistered})
			fmt.Println(PlayerData.PlayerId, PlayerData.Name, PlayerData.TimeRegistered)
		default:
			panic(err)
		}
		// get any error encountered during iteration
		err = rows.Err()
	}
	return Players
}

func InsertNewPlayer(playerId string, playerName string, TimeRegistered time.Time) (bool, error) {
	fmt.Println("Executing INSERT: InsertNewPlayer")

	sqlStatement := "INSERT INTO players (`PlayerId`,`Name`,`TimeRegistered`) VALUES (?,?,?)"

	stmt, err := db.Db.Prepare(sqlStatement)
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
