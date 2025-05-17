package handlers

import (
	"context"
	"github.com/artarts36/lowboard/registry/internal/model"
	"github.com/artarts36/lowboard/registry/internal/port/handlers/adapters"

	"github.com/artarts36/lowboard/registry/internal/port/generated/api"
)

func (h *Service) PageComponentsPost(ctx context.Context, req *api.PageComponentCreate) (*api.PageComponent, error) {
	return createAndAdapt(ctx, func(ctx context.Context) (*model.PageComponent, error) {
		return h.repo.PageComponent.Create(ctx, adapters.PageComponentFromCreateRequest(req))
	}, adapters.PageComponentToHTTP)
}
