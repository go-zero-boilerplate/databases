package sql_statement

import (
	"github.com/go-zero-boilerplate/databases"
	"github.com/go-zero-boilerplate/databases/sql"
)

type DeleteBuilder interface {
	OnDeleted(onDeleted OnDeleted) DeleteBuilder
	Where(condition string, args ...interface{}) DeleteBuilder
	WhereId(idFieldName string, id int64) DeleteBuilder
	Build() DeleteStatement
}

func NewDeleteBuilder(
	dialect databases.Dialect,
	db databases.Database,
	tableName string) DeleteBuilder {

	return &deleteBuilder{
		d: &deleteStatement{
			dialect:   dialect,
			db:        db,
			TableName: tableName,
		},
	}
}

type deleteBuilder struct {
	d *deleteStatement
}

func (d *deleteBuilder) OnDeleted(onDeleted OnDeleted) DeleteBuilder {
	d.d.onDeleted = onDeleted
	return d
}

func (d *deleteBuilder) Where(condition string, args ...interface{}) DeleteBuilder {
	d.d.Wheres = append(d.d.Wheres, &sql.WhereCondition{Condition: condition, Args: args})
	return d
}

func (d *deleteBuilder) WhereId(idFieldName string, id int64) DeleteBuilder {
	return d.Where(idFieldName+" = ?", id)
}

func (d *deleteBuilder) Build() DeleteStatement {
	return d.d
}
