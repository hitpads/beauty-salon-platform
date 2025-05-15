package domain

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
}
