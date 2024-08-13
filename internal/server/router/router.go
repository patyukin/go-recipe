package router

import (
	"context"
	"github.com/valyala/fasthttp"
	"go-recipe/internal/handler"
)

func Init(ctx context.Context, h *handler.Handler) func(requestCtx *fasthttp.RequestCtx) {
	r := func(requestCtx *fasthttp.RequestCtx) {
		switch string(requestCtx.Path()) {
		case "/authors":
			switch string(requestCtx.Method()) {
			case "POST":
				h.CreateAuthor(ctx, requestCtx)
			default:
				requestCtx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
			}
		case "/recipes":
			switch string(requestCtx.Method()) {
			case "GET":
				h.ListRecipes(ctx, requestCtx)
			case "POST":
				h.CreateRecipe(ctx, requestCtx)
			default:
				requestCtx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
			}
		default:
			if len(requestCtx.Path()) > len("/recipes/") && string(requestCtx.Path()[:9]) == "/recipes/" {
				switch string(requestCtx.Method()) {
				case "GET":
					h.GetRecipeByAuthor(ctx, requestCtx)
				case "PUT":
					h.UpdateRecipe(requestCtx)
				case "DELETE":
					h.DeleteRecipe(requestCtx)
				default:
					requestCtx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
				}
			} else {
				requestCtx.SetStatusCode(fasthttp.StatusNotFound)
			}
		}
	}

	return r
}
