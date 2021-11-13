package models

type User struct {
	ID       string `bun:"default:gen_random_uuid()"`
	Username string
	Password string
	Lift     []*LiftSession `bun:"rel:has-many,join:id=user_id"`
}
