package pg

import (
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

func (q *linksQ) Get() (*data.Link, error) {
	var result data.Link
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *linksQ) Insert(value data.Link) (*data.Link, error) {
	clauses := structs.Map(value)

	var result data.Link
	stmt := sq.Insert(linksTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return &result, err
}

func (q *linksQ) Delete(id string) error {
	stmt := sq.Delete(linksTableName).Where(sq.Eq{"id": id})
	err := q.db.Exec(stmt)
	return err
}

func (q *linksQ) FilterByID(ids ...string) data.LinksQ {
	q.sql = q.sql.Where(sq.Eq{"n.id": ids})
	return q
}

func (q *linksQ) Transaction(fn func(q data.LinksQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}
