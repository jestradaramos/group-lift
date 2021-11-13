package models

import "time"

type LiftSession struct {
	Date time.Time
	ID   string  `bun:"default:gen_random_uuid()"`
	Lift []*Lift `bun:"rel:has-many,join:id=`
}

type Lift struct {
	SessionID string
	Lift      string
	Weight    int
	Feel      string
}
