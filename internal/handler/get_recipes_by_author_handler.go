package handler

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

func (h *Handler) GetRecipeByAuthor(ctx context.Context, requestCtx *fasthttp.RequestCtx) {
	authorID := requestCtx.UserValue("author_id").(string)
	limit, offset := getPaginationParams(requestCtx)

	result, err := h.uc.GetRecipesByAuthorUseCase(ctx, authorID, limit, offset)
	if err != nil {
		log.Error().Msgf("failed GetRecipesByAuthorUseCase: %v", err)
		writeErrorResponse(requestCtx, fasthttp.StatusInternalServerError, err.Error())
		return
	}

	writeJsonResponse(requestCtx, fasthttp.StatusOK, result)
}
