package adapters

import (
	"database/sql"

	"github.com/artarts36/lowboard/registry/internal/port/generated/api"
)

func OptStringFromNullString(str *sql.NullString) api.OptString {
	if str.Valid {
		return api.NewOptString(str.String)
	}
	return api.OptString{}
}

func OptTimeFromNullTime(t sql.NullTime) api.OptDateTime {
	if t.Valid {
		return api.NewOptDateTime(t.Time)
	}
	return api.OptDateTime{}
}
