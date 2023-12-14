package data

import (
	"encoding/json"
	"time"
)

type LinksQ interface {
	New() LinksQ

	Get() (*Link, error)

	Insert(data Link) (*Link, error)

	FilterByID(id ...string) LinksQ
}

type Link struct {
	ID        string          `db:"id" structs:"-"`
	CreatedAt time.Time       `db:"created_at" structs:"created_at"`
	Value     json.RawMessage `db:"value" structs:"-"`
	Path      string          `db:"path" structs:"-"`
}
