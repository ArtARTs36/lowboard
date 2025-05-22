package model

type API struct {
	timestamps

	ID   string `db:"id"`
	Path string `db:"path"`
}

type APIAction struct {
	timestamps

	Name        string `db:"name" goqu:"skipUpdate"`
	Method      string `db:"method"`
	ApiID       string `db:"api_id"`
	Path        string `db:"path"`
	Description string `db:"description"`
}
