package models

import "time"

type LiftSession struct {
	Date   time.Time
	ID     string `bun:"default:gen_random_uuid()"`
	UserID string
	Lift   []*Lift `bun:"rel:has-many,join:id=session_id"`
}

type Lift struct {
	SessionID string
	Lift      string
	Weight    int
	Feel      string
}
