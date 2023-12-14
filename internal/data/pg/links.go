package pg

import (
	"database/sql"
	"fmt"
	"time"

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

func (q *linksQ) Get() (*data.Link, error) {
	var result data.Link
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *linksQ) Insert(value data.Link) (*data.Link, error) {
	value.CreatedAt = time.Now().UTC()
	clauses := structs.Map(value)
	clauses["value"] = value.Value
	clauses["path"] = value.Path
	clauses["id"] = value.ID

	var result data.Link
	stmt := sq.Insert(linksTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return &result, err
}

func (q *linksQ) FilterByID(ids ...string) data.LinksQ {
	q.sql = q.sql.Where(sq.Eq{"n.id": ids})
	return q
}
