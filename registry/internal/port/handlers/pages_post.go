package handlers

import (
	"context"
	"github.com/artarts36/lowboard/registry/internal/model"

	"github.com/artarts36/lowboard/registry/internal/port/generated/api"
	"github.com/artarts36/lowboard/registry/internal/port/handlers/adapters"
)

func (h *Service) PagesPost(ctx context.Context, req *api.PageCreate) (*api.Page, error) {
	return createAndAdapt(ctx, func(ctx context.Context) (*model.Page, error) {
		return h.repo.Page.Create(ctx, adapters.PageFromCreateRequest(req))
	}, adapters.PageToHTTP)
}
