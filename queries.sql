-- name: GetWizard :one
SELECT * FROM wizards WHERE id = $1 LIMIT 1;

-- name: GetAllWizards :many
SELECT * FROM wizards;

-- name: DeleteWizard :exec
DELETE FROM wizards WHERE id = $1;

-- name: CreateWizard :exec
INSERT INTO wizards (name, bio) VALUES ($1, $2) RETURNING *;