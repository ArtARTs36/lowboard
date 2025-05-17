package model

type API struct {
	timestamps

	ID   string `db:"id"`
	Path string `db:"path"`
}

type APIMethod struct {
	timestamps

	Name        string `db:"name" goqu:"skipUpdate"`
	ApiID       string `db:"api_id"`
	Path        string `db:"path"`
	Description string `db:"description"`
}
