package sql_statement

import (
	"github.com/go-zero-boilerplate/databases"
	"github.com/go-zero-boilerplate/databases/sql"
)

type InsertBuilder interface {
	OnInserted(onInserted OnInserted) InsertBuilder
	OnInsertedWithId(onInsertedWithId OnInsertedWithId) InsertBuilder
	Set(name string, value interface{}) InsertBuilder
	Build() InsertStatement
}

func NewInsertBuilder(
	dialect databases.Dialect,
	db databases.Database,
	tableName string) InsertBuilder {

	return &insertBuilder{
		i: &insertStatement{
			dialect:   dialect,
			db:        db,
			TableName: tableName,
		},
	}
}

type insertBuilder struct {
	i *insertStatement
}

func (i *insertBuilder) OnInserted(onInserted OnInserted) InsertBuilder {
	i.i.onInserted = onInserted
	return i
}

func (i *insertBuilder) OnInsertedWithId(onInsertedWithId OnInsertedWithId) InsertBuilder {
	i.i.onInsertedWithId = onInsertedWithId
	return i
}

func (i *insertBuilder) Set(name string, value interface{}) InsertBuilder {
	i.i.Columns = append(i.i.Columns, &sql.ColumnNameAndValue{Name: name, Value: value})
	return i
}

func (i *insertBuilder) Build() InsertStatement {
	return i.i
}
