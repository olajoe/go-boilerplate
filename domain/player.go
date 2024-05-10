package domain

import "time"

type Player struct {
	ID            string     `json:"ID"`
	FirstName     string     `json:"firstName"`
	LastName      string     `json:"lastName"`
	Position      string     `json:"position"`
	PreferredFoot string     `json:"preferredFoot"`
	UpdatedAt     *time.Time `json:"updated_at"`
	CreatedAt     *time.Time `json:"created_at"`
}
