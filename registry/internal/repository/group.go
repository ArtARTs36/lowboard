package repository

import "github.com/jmoiron/sqlx"

type Group struct {
	Page          *PageRepository
	PageComponent *PageComponentRepository
	Component     *ComponentRepository
	API           *APIRepository
	APIMethod     *APIMethodRepository
	Sidebar       *SidebarRepository
	SidebarLink   *SidebarLinkRepository
}

func NewGroup(db *sqlx.DB) *Group {
	return &Group{
		Page:          NewPageRepository(db),
		PageComponent: NewPageComponentRepository(db),
		Component:     NewComponentRepository(db),
		API:           NewAPIRepository(db),
		APIMethod:     NewAPIMethodRepository(db),
		Sidebar:       NewSidebarRepository(db),
		SidebarLink:   NewSidebarLinkRepository(db),
	}
}
