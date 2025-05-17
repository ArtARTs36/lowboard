package handlers

import (
	"context"

	"github.com/artarts36/lowboard/registry/internal/port/generated/api"
)

func (h *Service) DeletePageComponent(ctx context.Context, params api.DeletePageComponentParams) error {
	return h.repo.PageComponent.Delete(ctx, params.ComponentId)
}
