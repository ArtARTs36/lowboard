package model

import (
	"database/sql"
	"time"
)

type timestamps struct {
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
