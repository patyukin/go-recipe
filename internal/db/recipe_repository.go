package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	genapi "go-recipe/internal/gateway/http/gen"
)

func (r *Repository) SelectRecipesByAuthorID(ctx context.Context, authorID string, limit, offset int) ([]genapi.Recipe, error) {
	query := `SELECT r.id, r.title, r.instructions, r.created_at
		FROM recipes r
		JOIN recipe_authors ra ON r.id = ra.recipe_id
		WHERE ra.author_id = $1
		LIMIT $2 OFFSET $3`

	rows, err := r.db.QueryContext(ctx, query, authorID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed query: %v", err)
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			log.Error().Msgf("failed close rows: %v", err)
		}
	}(rows)

	var recipes []genapi.Recipe
	for rows.Next() {
		var recipe genapi.Recipe
		if err = rows.Scan(&recipe.Id, &recipe.Title, &recipe.Instructions, &recipe.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed scan: %v", err)
		}

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (r *Repository) CountRecipesByAuthorID(ctx context.Context, authorID string) (int, error) {
	query := `SELECT count(*) FROM recipes r
		JOIN recipe_authors ra ON r.id = ra.recipe_id
		WHERE ra.author_id = $1`

	row := r.db.QueryRowContext(ctx, query, authorID)

	var count int
	if err := row.Scan(&count); err != nil {
		return 0, fmt.Errorf("failed scan: %v", err)
	}

	return count, nil
}

func (r *Repository) SelectAllRecipes(ctx context.Context, limit, offset int) ([]genapi.Recipe, error) {
	query := `SELECT r.id, r.title, r.instructions, r.created_at
		FROM recipes r
		LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed query: %v", err)
	}

	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			log.Error().Msgf("failed close rows: %v", err)
		}
	}(rows)

	var recipes []genapi.Recipe
	for rows.Next() {
		var recipe genapi.Recipe
		if err = rows.Scan(&recipe.Id, &recipe.Title, &recipe.Instructions, &recipe.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed scan: %v", err)
		}

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (r *Repository) SelectCountRecipes(ctx context.Context) (int, error) {
	query := `SELECT count(*) FROM recipes`

	row := r.db.QueryRowContext(ctx, query)

	var count int
	if err := row.Scan(&count); err != nil {
		return 0, fmt.Errorf("failed scan: %v", err)
	}

	return count, nil
}
