package handlers

import (
	"context"
	"fmt"
	"github.com/artarts36/lowboard/registry/internal/model"
	"github.com/artarts36/lowboard/registry/internal/port/generated/api"
	"github.com/artarts36/lowboard/registry/internal/port/handlers/adapters"
)

func (h *Service) GetDefinition(ctx context.Context) (*api.Definition, error) {
	pages, err := h.repo.Page.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list pages: %w", err)
	}

	pageComponents, err := h.repo.PageComponent.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list page components: %w", err)
	}

	components, err := h.repo.Component.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list components: %w", err)
	}

	apis, err := h.repo.API.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list apis: %w", err)
	}

	apiActions, err := h.repo.APIMethod.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list api actions: %w", err)
	}

	apiActionMap := map[string]map[string]api.DefinitionAPIAction{}
	for _, action := range apiActions {
		if _, exists := apiActionMap[action.ApiID]; !exists {
			apiActionMap[action.ApiID] = map[string]api.DefinitionAPIAction{}
		}

		apiActionMap[action.ApiID][action.Name] = api.DefinitionAPIAction{
			Name:   action.Name,
			Path:   action.Path,
			Method: action.Method,
		}
	}

	pageComponentsMap := map[string][]api.DefinitionPageComponent{}

	for _, component := range pageComponents {
		if _, exists := pageComponentsMap[component.PageName]; !exists {
			pageComponentsMap[component.PageName] = []api.DefinitionPageComponent{}
		}

		pageComponentsMap[component.PageName] = append(pageComponentsMap[component.PageName], api.DefinitionPageComponent{
			ID:                component.ID,
			BaseComponentName: component.BaseComponentName,
			Config:            api.Map(component.Config),
		})
	}

	definition := &api.Definition{
		Pages:      make([]api.DefinitionPage, len(pages)),
		Components: map[string]api.DefinitionComponent{},
		Apis:       map[string]api.DefinitionAPI{},
	}

	for _, component := range components {
		definition.Components[component.Name] = api.DefinitionComponent{
			Name: component.Name,
		}
	}

	for i, page := range pages {
		definition.Pages[i] = api.DefinitionPage{
			Name:       page.Name,
			Title:      page.Title,
			Path:       page.Path,
			Components: pageComponentsMap[page.Name],
		}
	}

	for _, ap := range apis {
		definition.Apis[ap.ID] = api.DefinitionAPI{
			ID:      ap.ID,
			Path:    ap.Path,
			Actions: apiActionMap[ap.ID],
		}
	}

	sidebars, err := h.mapDefinitionsSidebars(ctx)
	if err != nil {
		return nil, err
	}

	definition.Sidebars = sidebars

	return definition, nil
}

func (h *Service) mapDefinitionsSidebars(ctx context.Context) (api.DefinitionSidebars, error) {
	res := api.DefinitionSidebars{}

	links, err := h.repo.SidebarLink.Tree(ctx)
	if err != nil {
		return nil, fmt.Errorf("list sidebar links: %w", err)
	}

	var wrapLink func(link *model.SidebarLink) api.DefinitionSidebarLink

	wrapLink = func(link *model.SidebarLink) api.DefinitionSidebarLink {
		l := api.DefinitionSidebarLink{
			Title:    link.Title,
			PageName: link.PageName,
			Children: make([]api.DefinitionSidebarLink, len(link.Children)),
			Icon:     adapters.OptStringFromNullString(link.Icon),
		}

		for i, child := range link.Children {
			l.Children[i] = wrapLink(child)
		}

		return l
	}

	for _, link := range links {
		if _, exists := res[link.SidebarName]; !exists {
			res[link.SidebarName] = api.DefinitionSidebar{
				Name:  link.SidebarName,
				Links: []api.DefinitionSidebarLink{},
			}
		}

		sb := res[link.SidebarName]

		sb.Links = append(sb.Links, wrapLink(link))

		res[link.SidebarName] = sb
	}

	return res, nil
}
