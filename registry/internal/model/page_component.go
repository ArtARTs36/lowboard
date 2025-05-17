package model

type PageComponent struct {
	timestamps

	ID string `db:"id" goqu:"skipUpdate"`

	PageName          string `db:"page_name"`
	BaseComponentName string `db:"base_component_name"`
	Config            JSONB  `db:"config"`
}
