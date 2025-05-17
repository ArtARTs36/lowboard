package model

type Page struct {
	timestamps

	Name  string `db:"name" goqu:"skipUpdate"`
	Path  string `db:"path"`
	Title string `db:"title"`
}
