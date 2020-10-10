package db

import "time"

// Game Properties
type Game struct {
	GameId  string
	Name    string
	Version float32
}

//Player Properties
type Player struct {
	PlayerId       string
	Name           string
	TimeRegistered time.Time
}

//Session Properties
type Session struct {
	SessionId      string
	PlayerId       string
	GameId         string
	TimeSessionEnd time.Time
}

//Session Rating Properties
type SessionRating struct {
	SessionId     string
	PlayerId      string
	Rating        int
	Comment       string
	TimeSubmitted time.Time
}
