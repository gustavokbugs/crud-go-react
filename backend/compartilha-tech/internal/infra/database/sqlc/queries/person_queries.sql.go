// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: person_queries.sql

package queries

import (
	"context"
	"database/sql"
	"time"
)

const deletePerson = `-- name: DeletePerson :exec
DELETE FROM person WHERE id = $1
`

func (q *Queries) DeletePerson(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deletePerson, id)
	return err
}

const getPersonById = `-- name: GetPersonById :one
SELECT id, name, age, active, created_at, updated_at FROM person WHERE id = $1
`

func (q *Queries) GetPersonById(ctx context.Context, id string) (Person, error) {
	row := q.db.QueryRowContext(ctx, getPersonById, id)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Age,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPersons = `-- name: GetPersons :many
SELECT id, name, age, active, created_at, updated_at FROM person
`

func (q *Queries) GetPersons(ctx context.Context) ([]Person, error) {
	rows, err := q.db.QueryContext(ctx, getPersons)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Person
	for rows.Next() {
		var i Person
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Age,
			&i.Active,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertPerson = `-- name: InsertPerson :exec
INSERT INTO person(id,name,age,active,created_at,updated_at)
VALUES($1,$2,$3,$4,$5,$6)
`

type InsertPersonParams struct {
	ID        string
	Name      string
	Age       sql.NullInt32
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) InsertPerson(ctx context.Context, arg InsertPersonParams) error {
	_, err := q.db.ExecContext(ctx, insertPerson,
		arg.ID,
		arg.Name,
		arg.Age,
		arg.Active,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const updatePerson = `-- name: UpdatePerson :exec
UPDATE person SET
name = COALESCE($2, name),
age = COALESCE($3, age),
active = COALESCE($4, active),
updated_at = $5
WHERE id = $1
`

type UpdatePersonParams struct {
	ID        string
	Name      sql.NullString
	Age       sql.NullInt32
	Active    sql.NullBool
	UpdatedAt time.Time
}

func (q *Queries) UpdatePerson(ctx context.Context, arg UpdatePersonParams) error {
	_, err := q.db.ExecContext(ctx, updatePerson,
		arg.ID,
		arg.Name,
		arg.Age,
		arg.Active,
		arg.UpdatedAt,
	)
	return err
}