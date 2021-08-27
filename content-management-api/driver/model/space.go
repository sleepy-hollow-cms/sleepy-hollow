package model

import "time"

type Space struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
