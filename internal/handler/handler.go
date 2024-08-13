package handler

import (
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	genapi "go-recipe/internal/gateway/http/gen"
	"strconv"
)

type UseCase interface {
	CreateAuthorUseCase(ctx context.Context, req genapi.CreateAuthorRequest) (genapi.PostAuthorsResponseObject, error)
	GetRecipesByAuthorUseCase(ctx context.Context, authorID string, limit, offset int) (genapi.GetRecipesResponseObject, error)
	GetRecipesUseCase(ctx context.Context, limit, offset int) (genapi.GetRecipesResponseObject, error)
}

type Handler struct {
	uc UseCase
}

func New(uc UseCase) *Handler {
	return &Handler{uc: uc}
}

func writeJsonResponse(ctx *fasthttp.RequestCtx, statusCode int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		writeErrorResponse(ctx, fasthttp.StatusInternalServerError, "Failed to marshal response")
		return
	}

	ctx.SetStatusCode(statusCode)
	ctx.SetContentType("application/json")
	ctx.SetBody(response)
}

func writeErrorResponse(ctx *fasthttp.RequestCtx, statusCode int, message string) {
	errorResponse := map[string]string{
		"error": message,
	}
	writeJsonResponse(ctx, statusCode, errorResponse)
}

func getPaginationParams(ctx *fasthttp.RequestCtx) (limit, offset int) {
	limit = 10
	offset = 0

	if l := string(ctx.QueryArgs().Peek("limit")); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v <= 10 {
			limit = v
		}
	}

	if o := string(ctx.QueryArgs().Peek("offset")); o != "" {
		if v, err := strconv.Atoi(o); err == nil {
			offset = v
		}
	}

	return limit, offset
}

func (h *Handler) UpdateRecipe(requestCtx *fasthttp.RequestCtx) {
	requestCtx.SetStatusCode(fasthttp.StatusOK)
}

func (h *Handler) DeleteRecipe(requestCtx *fasthttp.RequestCtx) {
	requestCtx.SetStatusCode(fasthttp.StatusOK)
}
