package store

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
)

type PostsStorage struct {
	db *sql.DB
}

type Posts struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    string   `json:"user_id"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	Tags      []string `json:"tags"`
}

func (s *PostsStorage) Create(ctx context.Context, post *Posts) error {
	query := `INSERT INTO posts (content,title,user_id,tags)
	VALUES ($1,$2,$3,$4) RETURNING id,created_at,updated_at
	`
	err := s.db.QueryRowContext(ctx, query, post.Content, post.Title, post.UserID, post.Tags, pq.Array(post.Tags)).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt)

	if err != nil {
		return err
	}
	return nil
}
