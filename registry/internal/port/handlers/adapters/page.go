package adapters

import (
	"github.com/artarts36/lowboard/registry/internal/model"
	"github.com/artarts36/lowboard/registry/internal/port/generated/api"
)

func PageToHTTP(page *model.Page) api.Page {
	return api.Page{
		Name:      page.Name,
		Title:     page.Title,
		Path:      page.Path,
		CreatedAt: page.CreatedAt,
		UpdatedAt: OptTimeFromNullTime(page.UpdatedAt),
	}
}

func PagesToHTTP(pages []*model.Page) []api.Page {
	ps := make([]api.Page, len(pages))

	for i, page := range pages {
		ps[i] = PageToHTTP(page)
	}

	return ps
}

func PageFromCreateRequest(req *api.PageCreate) *model.Page {
	return &model.Page{
		Name:  req.Name,
		Path:  req.Path,
		Title: req.Title,
	}
}

func PageFromUpdateRequest(req *api.PageUpdate) *model.Page {
	return &model.Page{
		Name:  req.Name,
		Path:  req.Path,
		Title: req.Title,
	}
}
