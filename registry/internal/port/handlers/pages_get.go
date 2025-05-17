package handlers

import (
	"context"
	"github.com/artarts36/lowboard/registry/internal/port/handlers/adapters"

	"github.com/artarts36/lowboard/registry/internal/port/generated/api"
)

func (h *Service) PagesGet(ctx context.Context) ([]api.Page, error) {
	return listAndAdapt(ctx, h.repo.Page.List, adapters.PagesToHTTP)
}
