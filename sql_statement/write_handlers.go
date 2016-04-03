package sql_statement

type OnInserted interface {
	OnInserted()
}

type OnInsertedWithId interface {
	OnInsertedWithId(id int64)
}

type OnDeleted interface {
	OnDeleted()
}

type OnUpdated interface {
	OnUpdated()
}
