-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (
    username,
    email,
    secret_code
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: SetUsedVerifyEmail :one
UPDATE verify_emails SET is_used = true
WHERE id = $1 AND secret_code = $2
RETURNING *;