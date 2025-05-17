package handlers

import (
	"context"
	"github.com/artarts36/lowboard/registry/internal/model"
	"github.com/artarts36/lowboard/registry/internal/port/handlers/adapters"

	"github.com/artarts36/lowboard/registry/internal/port/generated/api"
)

func (h *Service) PagesPageNamePut(ctx context.Context, req *api.PageUpdate, params api.PagesPageNamePutParams) (*api.Page, error) {
	return updateAndAdapt(ctx, func(ctx context.Context) (*model.Page, error) {
		return h.repo.Page.Update(ctx, adapters.PageFromUpdateRequest(req), params.PageName)
	}, adapters.PageToHTTP)
}
