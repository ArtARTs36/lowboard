package model

import "database/sql"

type Sidebar struct {
	timestamps

	Name string `db:"sidebar"`
}

type SidebarLink struct {
	timestamps

	ID string `db:"id"` // uuid

	SidebarName string `db:"sidebar_name"`
	PageName    string `db:"page_name"`
	Title       string `db:"title"`

	ParentID sql.NullString `db:"parent_id"`

	Children []*SidebarLink `db:"-"`
}
