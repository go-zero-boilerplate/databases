package sql_statement

import (
	"github.com/thcyron/sqlbuilder"

	"github.com/go-zero-boilerplate/databases"
	"github.com/go-zero-boilerplate/databases/sql"
)

type DeleteStatement interface {
	BaseStatement
}

type deleteStatement struct {
	dialect databases.Dialect
	db      databases.Database

	onDeleted OnDeleted

	TableName string
	Wheres    []*sql.WhereCondition
}

func (d *deleteStatement) Execute() error {
	builder := sqlbuilder.Delete().Dialect(d.dialect).From(d.TableName)

	for _, w := range d.Wheres {
		builder = builder.Where(w.Condition, w.Args...)
	}

	query, args := builder.Build()

	_, err := d.db.Exec(query, args...)
	if err != nil {
		return err
	}

	if d.onDeleted != nil {
		d.onDeleted.OnDeleted()
	}

	return nil
}
