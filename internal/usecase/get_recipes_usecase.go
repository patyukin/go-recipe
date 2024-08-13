package usecase

import (
	"context"
	"fmt"
	genapi "go-recipe/internal/gateway/http/gen"
)

func (uc *UseCase) GetRecipesUseCase(ctx context.Context, limit, offset int) (genapi.GetRecipesResponseObject, error) {
	res, err := uc.registry.GetRepo().SelectAllRecipes(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed GetRecipesByAuthorUseCase: %w", err)
	}

	cnt, err := uc.registry.GetRepo().SelectCountRecipes(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed GetRecipesByAuthorUseCase: %w", err)
	}

	return genapi.GetRecipes200JSONResponse{Recipes: &res, Total: &cnt}, nil
}
