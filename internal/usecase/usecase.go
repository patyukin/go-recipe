package usecase

import (
	"go-recipe/internal/client"
	"go-recipe/internal/db"
)

type UseCase struct {
	registry   *db.Client
	httpClient *client.Client
}

func New(registry *db.Client, httpClient *client.Client) *UseCase {
	return &UseCase{
		registry:   registry,
		httpClient: httpClient,
	}
}
