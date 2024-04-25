// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: videos.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const getAllVideos = `-- name: GetAllVideos :many
SELECT id, created_at, updated_at, title, description, image_url, authors, published_at, url, view_count, star_rating, star_count, channel_id FROM videos
`

func (q *Queries) GetAllVideos(ctx context.Context) ([]Video, error) {
	rows, err := q.db.QueryContext(ctx, getAllVideos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Video
	for rows.Next() {
		var i Video
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Description,
			&i.ImageUrl,
			&i.Authors,
			&i.PublishedAt,
			&i.Url,
			&i.ViewCount,
			&i.StarRating,
			&i.StarCount,
			&i.ChannelID,
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

const insertVideo = `-- name: InsertVideo :exec
INSERT INTO videos
  (id, created_at, updated_at, title, description, image_url, authors, published_at, url, view_count, star_rating, star_count, channel_id)
VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
`

type InsertVideoParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Description string
	ImageUrl    string
	Authors     string
	PublishedAt time.Time
	Url         string
	ViewCount   string
	StarRating  string
	StarCount   string
	ChannelID   string
}

func (q *Queries) InsertVideo(ctx context.Context, arg InsertVideoParams) error {
	_, err := q.db.ExecContext(ctx, insertVideo,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Description,
		arg.ImageUrl,
		arg.Authors,
		arg.PublishedAt,
		arg.Url,
		arg.ViewCount,
		arg.StarRating,
		arg.StarCount,
		arg.ChannelID,
	)
	return err
}