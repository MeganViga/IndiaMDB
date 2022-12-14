// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: rating.sql

package db

import (
	"context"
)

const getMovieListBasedOnRating = `-- name: GetMovieListBasedOnRating :many
SELECT id, movie_id, user_id, rating 
FROM "rating"
WHERE rating = $1
`

func (q *Queries) GetMovieListBasedOnRating(ctx context.Context, rating int32) ([]Rating, error) {
	rows, err := q.db.QueryContext(ctx, getMovieListBasedOnRating, rating)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Rating
	for rows.Next() {
		var i Rating
		if err := rows.Scan(
			&i.ID,
			&i.MovieID,
			&i.UserID,
			&i.Rating,
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

const getRatingsGivenByUser = `-- name: GetRatingsGivenByUser :many
SELECT id, movie_id, user_id, rating
FROM "rating"
WHERE user_id = $1
`

func (q *Queries) GetRatingsGivenByUser(ctx context.Context, userID int32) ([]Rating, error) {
	rows, err := q.db.QueryContext(ctx, getRatingsGivenByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Rating
	for rows.Next() {
		var i Rating
		if err := rows.Scan(
			&i.ID,
			&i.MovieID,
			&i.UserID,
			&i.Rating,
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

const getRatingsGivenToMovie = `-- name: GetRatingsGivenToMovie :many
SELECT id, movie_id, user_id, rating FROM "rating"
WHERE movie_id = $1
`

func (q *Queries) GetRatingsGivenToMovie(ctx context.Context, movieID int32) ([]Rating, error) {
	rows, err := q.db.QueryContext(ctx, getRatingsGivenToMovie, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Rating
	for rows.Next() {
		var i Rating
		if err := rows.Scan(
			&i.ID,
			&i.MovieID,
			&i.UserID,
			&i.Rating,
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

const giveRating = `-- name: GiveRating :one
INSERT INTO "rating"(
   movie_id,
   user_id,
   rating
) VALUES ($1,$2,$3) RETURNING id, movie_id, user_id, rating
`

type GiveRatingParams struct {
	MovieID int32 `json:"movie_id"`
	UserID  int32 `json:"user_id"`
	Rating  int32 `json:"rating"`
}

func (q *Queries) GiveRating(ctx context.Context, arg GiveRatingParams) (Rating, error) {
	row := q.db.QueryRowContext(ctx, giveRating, arg.MovieID, arg.UserID, arg.Rating)
	var i Rating
	err := row.Scan(
		&i.ID,
		&i.MovieID,
		&i.UserID,
		&i.Rating,
	)
	return i, err
}
