package domain

import "time"

type Audit struct {
	ID          string
	Title       string
	Norm        string
	CreatedAt   time.Time
	Status      string
	Assignments []Assignment
}
