package handler

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

func (h *Handler) ListRecipes(ctx context.Context, requestCtx *fasthttp.RequestCtx) {
	limit, offset := getPaginationParams(requestCtx)

	result, err := h.uc.GetRecipesUseCase(ctx, limit, offset)
	if err != nil {
		log.Error().Msgf("failed GetRecipesByAuthorUseCase: %v", err)
		writeErrorResponse(requestCtx, fasthttp.StatusInternalServerError, err.Error())
		return
	}

	writeJsonResponse(requestCtx, fasthttp.StatusOK, result)
}
