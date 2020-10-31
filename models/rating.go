package models

import "time"

//Session Rating Properties
type Rating struct {
	SessionId     string
	PlayerId      string
	Rating        int
	Comment       string
	TimeSubmitted time.Time
}

func (s Rating) IsEmpty() bool {
	return s.PlayerId == ""
}
