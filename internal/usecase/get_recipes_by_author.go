package usecase

import (
	"context"
	"fmt"
	genapi "go-recipe/internal/gateway/http/gen"
)

func (uc *UseCase) GetRecipesByAuthorUseCase(ctx context.Context, authorID string, limit, offset int) (genapi.GetRecipesResponseObject, error) {
	res, err := uc.registry.GetRepo().SelectRecipesByAuthorID(ctx, authorID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed GetRecipesByAuthorUseCase: %w", err)
	}

	cnt, err := uc.registry.GetRepo().CountRecipesByAuthorID(ctx, authorID)
	if err != nil {
		return nil, fmt.Errorf("failed GetRecipesByAuthorUseCase: %w", err)
	}

	return genapi.GetRecipes200JSONResponse{Recipes: &res, Total: &cnt}, nil
}
