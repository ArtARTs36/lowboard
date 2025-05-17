package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/go-faster/jx"
)

type JSONB map[string]jx.Raw

func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *JSONB) Scan(value interface{}) error {
	bs := []byte{}

	switch v := value.(type) {
	case []byte:
		bs = v
	case string:
		bs = []byte(v)
	default:
		return fmt.Errorf("expected []byte or string, got %T", value)
	}

	return json.Unmarshal(bs, &a)
}
