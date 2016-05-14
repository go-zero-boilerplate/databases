# databases


## Features

- Paginator - iterate through entities in the Database by using Offset and Limit in the background to "page" through your entries but in your code all you need is to call `Next()`
- Sql Builders (JOINS not yet supported)
- Dialect abstraction


## Code generator

For a sql model with code generator see [gensql](https://github.com/go-zero-boilerplate/gensql). It currently generates entities, iterators and repositories.

## Database method `DeferredDone`

This method could be used like this.

```
tx, err := db.BeginTx()
if err != nil {
    log.Fatal(err)
}

var err error
err = doSomething()
defer tx.DeferredDone(&err)
```

This code basically starts the transaction and then defers the `DeferredDone` which will based on the `err` determine if it should call `CommitTx` or `RollbackTx`.

## Acknowledgments

- https://github.com/thcyron/sqlbuilder