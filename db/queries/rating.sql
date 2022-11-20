-- name: GiveRating :one
INSERT INTO "rating"(
   movie_id,
   user_id,
   rating
) VALUES ($1,$2,$3) RETURNING *;

-- name: GetMovieListBasedOnRating :many
SELECT * 
FROM "rating"
WHERE rating = $1;

-- name: GetRatingsGivenByUser :many
SELECT *
FROM "rating"
WHERE user_id = $1;

-- name: GetRatingsGivenToMovie :many
SELECT * FROM "rating"
WHERE movie_id = $1;
