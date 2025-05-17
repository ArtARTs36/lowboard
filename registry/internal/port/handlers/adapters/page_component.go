package adapters

import (
	"github.com/artarts36/lowboard/registry/internal/model"
	"github.com/artarts36/lowboard/registry/internal/port/generated/api"
)

func PageComponentToHTTP(pc *model.PageComponent) api.PageComponent {
	return api.PageComponent{
		ID:                pc.ID,
		PageName:          pc.PageName,
		BaseComponentName: pc.BaseComponentName,
		Config:            api.PageComponentConfig(pc.Config),
		CreatedAt:         pc.CreatedAt,
		UpdatedAt:         OptTimeFromNullTime(pc.UpdatedAt),
	}
}

func PageComponentsToHTTP(pcs []*model.PageComponent) []api.PageComponent {
	res := make([]api.PageComponent, len(pcs))

	for i, pc := range pcs {
		res[i] = PageComponentToHTTP(pc)
	}

	return res
}

func PageComponentFromCreateRequest(req *api.PageComponentCreate) *model.PageComponent {
	return &model.PageComponent{
		PageName:          req.PageName,
		BaseComponentName: req.BaseComponentName,
		Config:            model.JSONB(req.Config),
	}
}

func PageComponentFromPutRequest(req *api.PageComponentUpdate, id string) *model.PageComponent {
	return &model.PageComponent{
		ID:                id,
		PageName:          req.PageName,
		BaseComponentName: req.BaseComponentName,
		Config:            model.JSONB(req.Config),
	}
}
