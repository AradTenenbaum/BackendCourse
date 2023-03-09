-- name: AddCurrency :one
INSERT INTO currency ("desc") VALUES (
  $1
) RETURNING *;