package models

import "time"

//Player Properties
type Player struct {
	PlayerId       string    `json:"playerId"`
	Name           string    `json:"name"`
	TimeRegistered time.Time `json:"timeRegistered"`
}
