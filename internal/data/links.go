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
	ID        string          `db:"id" structs:"id"`
	ExpiredAt time.Time       `db:"expired_at" structs:"expired_at"`
	Value     json.RawMessage `db:"value" structs:"value"`
	Path      string          `db:"path" structs:"path"`
}
