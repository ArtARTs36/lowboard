package handlers

import (
	"context"
	"github.com/artarts36/lowboard/registry/internal/model"
	"github.com/artarts36/lowboard/registry/internal/port/handlers/adapters"

	"github.com/artarts36/lowboard/registry/internal/port/generated/api"
)

func (h *Service) PagesPageNameGet(ctx context.Context, params api.PagesPageNameGetParams) (*api.Page, error) {
	return getAndAdapt(ctx, func(ctx context.Context) (*model.Page, error) {
		return h.repo.Page.Get(ctx, params.PageName)
	}, adapters.PageToHTTP)
}
