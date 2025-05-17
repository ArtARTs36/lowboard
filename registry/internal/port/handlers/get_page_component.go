package handlers

import (
	"context"
	"github.com/artarts36/lowboard/registry/internal/model"
	"github.com/artarts36/lowboard/registry/internal/port/handlers/adapters"

	"github.com/artarts36/lowboard/registry/internal/port/generated/api"
)

func (h *Service) GetPageComponent(ctx context.Context, params api.GetPageComponentParams) (*api.PageComponent, error) {
	return getAndAdapt(ctx, func(ctx context.Context) (*model.PageComponent, error) {
		return h.repo.PageComponent.Get(ctx, params.ComponentId)
	}, adapters.PageComponentToHTTP)
}
