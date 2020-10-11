package db

import "time"

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

func (s SessionRating) IsEmpty() bool {
	return s.PlayerId == ""
}
