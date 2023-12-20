package pg

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/dl-solarity/frontend-link-shortener-svc/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const linksTableName = "links"

type linksQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func NewLinksQ(db *pgdb.DB) data.LinksQ {
	return &linksQ{
		db:  db.Clone(),
		sql: sq.Select("n.*").From(fmt.Sprintf("%s as n", linksTableName)),
	}
}

func (q *linksQ) New() data.LinksQ {
	return NewLinksQ(q.db)
}

func (q *linksQ) Get(ctx context.Context, id string) (*data.Link, error) {
	var result data.Link
	err := q.db.GetContext(ctx, &result, q.sql.Where(sq.Eq{"id": id}))
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *linksQ) Insert(ctx context.Context, value data.Link) (*data.Link, error) {
	stmt := sq.Insert(linksTableName).
		SetMap(structs.Map(value)).
		Suffix("on conflict (id) do update set path = EXCLUDED.path, value = EXCLUDED.value, expired_at = EXCLUDED.expired_at").
		Suffix("returning *")

	var result data.Link
	err := q.db.GetContext(ctx, &result, stmt)

	return &result, err
}
