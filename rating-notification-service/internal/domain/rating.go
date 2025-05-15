package domain

import "time"

type Rating struct {
	ID        string
	MasterID  string
	UserID    string
	Score     int
	Comment   string
	CreatedAt time.Time
}
