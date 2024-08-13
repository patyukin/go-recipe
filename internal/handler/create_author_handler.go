package handler

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
	genapi "go-recipe/internal/gateway/http/gen"
)

func (h *Handler) CreateAuthor(ctx context.Context, requestCtx *fasthttp.RequestCtx) {
	log.Info().Msgf("create author")
	var req genapi.CreateAuthorRequest
	if err := json.Unmarshal(requestCtx.PostBody(), &req); err != nil {
		log.Error().Msgf("create author: %v", err)
		requestCtx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	log.Info().Msgf("create author: %v", req)
	result, err := h.uc.CreateAuthorUseCase(ctx, req)
	if err != nil {
		log.Info().Msgf("create author: %v", err)

		requestCtx.SetStatusCode(fasthttp.StatusInternalServerError)
		response, errResult := json.Marshal(result)
		if errResult != nil {
			requestCtx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}

		requestCtx.SetStatusCode(fasthttp.StatusOK)
		requestCtx.SetBody(response)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		requestCtx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	requestCtx.SetContentType("application/json")
	requestCtx.SetStatusCode(fasthttp.StatusCreated)
	requestCtx.SetBody(response)
}
