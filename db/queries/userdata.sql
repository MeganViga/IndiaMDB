-- name: CreateUser :one
INSERT INTO "userdata"(
    name,
    email,
    role,
    passwordhash
) VALUES ($1,$2,$3,$4) RETURNING *;

-- name: ResetPassword :one
UPDATE "userdata" SET passwordhash = $1 WHERE id = $2 RETURNING name,email;

-- name: GetUserByEmail :one
SELECT *
FROM "userdata"
WHERE email = $1;

-- name: GetUserByName :one
SELECT *
FROM "userdata"
WHERE name = $1;