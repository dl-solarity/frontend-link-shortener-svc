package data

import (
	"encoding/json"
	"time"
)

type LinksQ interface {
	New() LinksQ

	Get() (*Link, error)

	Insert(data Link) (*Link, error)

	Delete(id string) error

	FilterByID(id ...string) LinksQ

	Transaction(fn func(q LinksQ) error) error
}

type Link struct {
	ID        string          `db:"id" structs:"-"`
	CreatedAt time.Time       `db:"created_at" structs:"created_at"`
	Value     json.RawMessage `db:"value" structs:"-"`
	Path      string          `db:"path" structs:"-"`
}
