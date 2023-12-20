package data

import (
	"context"
	"encoding/json"
	"time"
)

type LinksQ interface {
	New() LinksQ

	Get(ctx context.Context, id string) (*Link, error)

	Insert(ctx context.Context, data Link) (*Link, error)
}

type Link struct {
	ID        string          `db:"id" structs:"id" json:"id"`
	ExpiredAt time.Time       `db:"expired_at" structs:"expired_at" json:"expired_at"`
	Value     json.RawMessage `db:"value" structs:"value" json:"value"`
	Path      string          `db:"path" structs:"path" json:"path"`
}
