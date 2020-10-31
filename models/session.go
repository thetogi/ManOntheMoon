package models

import "time"

//Session Properties
type Session struct {
	SessionId      string
	PlayerId       string
	TimeSessionEnd time.Time
}
