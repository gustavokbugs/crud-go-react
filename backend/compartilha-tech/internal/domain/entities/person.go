package entities

import (
	"time"

	"github.com/google/uuid"
)

type Person struct {
	ID        string
	Name      string
	Age       int
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPerson() *Person {
	return &Person{
		ID:        uuid.NewString(),
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
