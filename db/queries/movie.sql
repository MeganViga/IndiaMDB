-- name: CreateMovie :one
INSERT INTO "movie"(
    name,
    summary,
    language,
    genre,
    release_date
) VALUES ($1,$2,$3,$4,$5) RETURNING *;

-- name: EditMovie :many
UPDATE "movie"
SET name = $1,
    summary = $2,
    language = $3,
    genre = $4,
    release_date = $5
WHERE id = $6 RETURNING *;

-- name: GetMovieByName :one
SELECT *
FROM "movie"
WHERE name = $1;

-- name: GetMovieByID :one
SELECT * 
FROM "movie"
WHERE id = $1;

-- name: GetMoviesByLanguage :many
SELECT *
FROM "movie"
WHERE language = $1;

-- name: GetMoviesByGenre :many
SELECT *
FROM "movie"
WHERE genre = $1;