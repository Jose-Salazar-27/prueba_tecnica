package models

import "time"

type User struct {
	Id         int
	Name       string
	LastName   string
	Profession string
	Created_at time.Time
}
