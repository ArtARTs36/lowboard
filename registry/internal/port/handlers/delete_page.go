package handlers

import (
	"context"

	"github.com/artarts36/lowboard/registry/internal/port/generated/api"
)

func (h *Service) DeletePage(ctx context.Context, params api.DeletePageParams) error {
	return h.repo.Page.Delete(ctx, params.PageName)
}
