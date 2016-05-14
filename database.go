package databases

type Database interface {
	IsErrNoRows(err error) bool

	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (ExecResult, error)
	QueryRow(query string, args ...interface{}) ResultRow
	Query(query string, args ...interface{}) (ResultRows, error)

	BeginTx() (Database, error)
	CommitTx() error
	RollbackTx() error

	//DeferredDone should be called as a deferred method. The error pointer is used to detemine if should Commit or Rollback the transaction
	DeferredDone(err *error) error
}
