// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package queries

import (
	"database/sql"
	"time"
)

type Person struct {
	ID        string
	Name      string
	Age       sql.NullInt32
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
