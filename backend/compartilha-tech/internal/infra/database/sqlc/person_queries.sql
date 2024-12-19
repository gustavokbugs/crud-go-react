-- name: InsertPerson :exec
INSERT INTO person(id, name, age, active, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetPersons :many
SELECT * FROM person;

-- name: GetPersonById :one
SELECT * FROM person WHERE id = $1;

-- name: UpdatePerson :exec
UPDATE person
SET name = COALESCE($2, name),
    age = COALESCE($3, age),
    active = COALESCE($4, active),
    updated_at = $5
WHERE id = $1;

-- name: DeletePerson :exec
DELETE FROM person WHERE id = $1;
