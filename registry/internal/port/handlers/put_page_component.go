package handlers

import (
	"context"
	"github.com/artarts36/lowboard/registry/internal/model"
	"github.com/artarts36/lowboard/registry/internal/port/handlers/adapters"

	"github.com/artarts36/lowboard/registry/internal/port/generated/api"
)

func (h *Service) PutPageComponent(ctx context.Context, req *api.PageComponentUpdate, params api.PutPageComponentParams) (*api.PageComponent, error) {
	return updateAndAdapt(ctx, func(ctx context.Context) (*model.PageComponent, error) {
		return h.repo.PageComponent.Update(ctx, adapters.PageComponentFromPutRequest(req, params.ComponentId))
	}, adapters.PageComponentToHTTP)
}
