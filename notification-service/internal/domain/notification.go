package domain

import "time"

type Notification struct {
	ID        string
	UserID    string
	Message   string
	IsRead    bool
	CreatedAt time.Time
}
