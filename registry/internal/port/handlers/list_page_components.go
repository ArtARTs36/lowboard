package handlers

import (
	"context"
	"github.com/artarts36/lowboard/registry/internal/model"
	"github.com/artarts36/lowboard/registry/internal/port/handlers/adapters"

	"github.com/artarts36/lowboard/registry/internal/port/generated/api"
)

func (h *Service) ListPageComponents(ctx context.Context) ([]api.PageComponent, error) {
	return listAndAdapt(ctx, func(ctx context.Context) ([]*model.PageComponent, error) {
		return h.repo.PageComponent.List(ctx)
	}, adapters.PageComponentsToHTTP)
}
