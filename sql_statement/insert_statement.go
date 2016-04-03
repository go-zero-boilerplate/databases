package sql_statement

import (
	"github.com/thcyron/sqlbuilder"

	"github.com/go-zero-boilerplate/databases"
	"github.com/go-zero-boilerplate/databases/sql"
)

type InsertStatement interface {
	BaseStatement
}

type insertStatement struct {
	dialect databases.Dialect
	db      databases.Database

	onInserted       OnInserted
	onInsertedWithId OnInsertedWithId

	TableName string
	Columns   []*sql.ColumnNameAndValue
}

func (i *insertStatement) Execute() error {
	builder := sqlbuilder.Insert().Dialect(i.dialect).Into(i.TableName)

	for _, c := range i.Columns {
		builder = builder.Set(c.Name, c.Value)
	}

	query, args, _ := builder.Build()

	result, err := i.db.Exec(query, args...)
	if err != nil {
		return err
	}

	//TODO: Think this is possible for instance if int64 ID's are not used
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		if i.onInserted != nil {
			i.onInserted.OnInserted()
		}
		if i.onInsertedWithId != nil {
			//TODO: This seems to be unexpected, will never get fired
		}
	} else {
		if i.onInsertedWithId != nil {
			i.onInsertedWithId.OnInsertedWithId(lastInsertId)
		}
	}

	return nil
}
