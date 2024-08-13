package db

import (
	"context"
	"fmt"
	genapi "go-recipe/internal/gateway/http/gen"
)

func (r *Repository) InsertIntoAuthors(ctx context.Context, in genapi.CreateAuthorRequest, userUUID string) (string, string, error) {
	query := `INSERT INTO authors (user_id, bio) VALUES ($1, $2) RETURNING id, bio`

	var id string
	var bio string

	row := r.db.QueryRowContext(ctx, query, userUUID, in.Bio)
	if row.Err() != nil {
		return "", "", fmt.Errorf("failed to insert into authors: %w", row.Err())
	}

	err := row.Scan(&id, &bio)
	if err != nil {
		return "", "", fmt.Errorf("failed to scan row: %w", err)
	}

	return id, bio, nil
}
