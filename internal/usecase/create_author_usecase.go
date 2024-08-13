package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go-recipe/internal/db"
	genapi "go-recipe/internal/gateway/http/gen"
)

func (uc *UseCase) CreateAuthorUseCase(ctx context.Context, req genapi.CreateAuthorRequest) (genapi.PostAuthorsResponseObject, error) {
	var authorUUID uuid.UUID
	var authorID string
	var bio string
	var err error

	userUUID, errClient := uc.httpClient.CreateUser(req)
	if errClient != nil {
		return genapi.PostAuthors400JSONResponse{}, fmt.Errorf("failed to create user: %w", errClient)
	}

	err = uc.registry.ReadCommitted(ctx, func(ctx context.Context, repo *db.Repository) error {
		authorID, bio, err = repo.InsertIntoAuthors(ctx, req, userUUID)
		if err != nil {
			return fmt.Errorf("failed to create author: %w", err)
		}

		log.Info().Msgf("created author: %s, bio: %s", authorID, bio)

		return nil
	})
	if err != nil {
		return genapi.PostAuthors400JSONResponse{}, fmt.Errorf("failed to execute handler: %w", err)
	}

	authorUUID, err = uuid.Parse(authorID)
	if err != nil {
		return genapi.PostAuthors400JSONResponse{}, fmt.Errorf("failed to parse author id: %w", err)
	}

	return genapi.PostAuthors201JSONResponse{Id: authorUUID, Bio: &bio}, nil
}
