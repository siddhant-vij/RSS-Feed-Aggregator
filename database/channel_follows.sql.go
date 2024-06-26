// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: channel_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const deleteChannelFollow = `-- name: DeleteChannelFollow :exec
DELETE FROM channel_follows
WHERE user_id = $1 AND channel_id = $2
`

type DeleteChannelFollowParams struct {
	UserID    uuid.UUID
	ChannelID string
}

func (q *Queries) DeleteChannelFollow(ctx context.Context, arg DeleteChannelFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteChannelFollow, arg.UserID, arg.ChannelID)
	return err
}

const getAllChannelFollows = `-- name: GetAllChannelFollows :many
SELECT id, created_at, updated_at, user_id, channel_id FROM channel_follows
`

func (q *Queries) GetAllChannelFollows(ctx context.Context) ([]ChannelFollow, error) {
	rows, err := q.db.QueryContext(ctx, getAllChannelFollows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChannelFollow
	for rows.Next() {
		var i ChannelFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
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

const getChannelFollowsForUser = `-- name: GetChannelFollowsForUser :many
SELECT id, created_at, updated_at, user_id, channel_id FROM channel_follows
WHERE user_id = $1
`

func (q *Queries) GetChannelFollowsForUser(ctx context.Context, userID uuid.UUID) ([]ChannelFollow, error) {
	rows, err := q.db.QueryContext(ctx, getChannelFollowsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChannelFollow
	for rows.Next() {
		var i ChannelFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
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

const getOtherChannelsForUser = `-- name: GetOtherChannelsForUser :many
SELECT id, created_at, updated_at, name, url, last_fetched_at FROM channels
WHERE id NOT IN (
  SELECT channel_id
  FROM channel_follows
  WHERE user_id = $1
)
`

func (q *Queries) GetOtherChannelsForUser(ctx context.Context, userID uuid.UUID) ([]Channel, error) {
	rows, err := q.db.QueryContext(ctx, getOtherChannelsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Channel
	for rows.Next() {
		var i Channel
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Url,
			&i.LastFetchedAt,
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

const getUserFollowedChannels = `-- name: GetUserFollowedChannels :many
SELECT channels.id, channels.created_at, channels.updated_at, channels.name, channels.url, channels.last_fetched_at
FROM channel_follows
JOIN channels
ON channel_follows.channel_id = channels.id
WHERE channel_follows.user_id = $1
`

func (q *Queries) GetUserFollowedChannels(ctx context.Context, userID uuid.UUID) ([]Channel, error) {
	rows, err := q.db.QueryContext(ctx, getUserFollowedChannels, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Channel
	for rows.Next() {
		var i Channel
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Url,
			&i.LastFetchedAt,
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

const insertChannelFollow = `-- name: InsertChannelFollow :exec
INSERT INTO channel_follows
  (id, created_at, updated_at, user_id, channel_id)
VALUES
  ($1, $2, $3, $4, $5)
`

type InsertChannelFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	ChannelID string
}

func (q *Queries) InsertChannelFollow(ctx context.Context, arg InsertChannelFollowParams) error {
	_, err := q.db.ExecContext(ctx, insertChannelFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.ChannelID,
	)
	return err
}
