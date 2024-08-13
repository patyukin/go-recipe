package handler

import (
	"context"
	"github.com/valyala/fasthttp"
)

func (h *Handler) CreateRecipe(_ context.Context, requestCtx *fasthttp.RequestCtx) {
	requestCtx.SetStatusCode(fasthttp.StatusCreated)
	requestCtx.SetBody([]byte("recipe created"))
}
