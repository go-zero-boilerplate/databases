# databases


## Features

- Paginator - iterate through entities in the Database by using Offset and Limit in the background to "page" through your entries but in your code all you need is to call `Next()`
- Sql Builders (JOINS not yet supported)
- Dialect abstraction


## Code generator

For a sql model with code generator see [gensql](https://github.com/go-zero-boilerplate/gensql). It currently generates entities, iterators and repositories.

## Usage of `DeferredRollbackIfNotHandled`

This method should be used with the `defer` keyword. This is an example:

```
tx, err := db.BeginTx()
if err != nil {
    log.Fatalf("Cannot begin transaction, error: %s", err.Error())
}

doSomething()

defer tx.DeferredRollbackIfNotHandled()
```

In this example if we did a Commit/Rollback as part of `doSomething`, the `DeferredRollbackIfNotHandled` would actually do nothing. If we however did not do anything, the transaction will automatically (at `defer` time) be rolled back.

Note that the `DeferredRollbackIfNotHandled` also returns an error (which would come from the Rollback), if we want to act upon that we could do this:

```
tx, err := db.BeginTx()
if err != nil {
    log.Fatalf("Cannot begin transaction, error: %s", err.Error())
}

doSomething()

defer func(){
    if err = tx.DeferredRollbackIfNotHandled(); err != nil {
        //Handle the deferred (rollback) error here
    }
}()
```

## Acknowledgments

- https://github.com/thcyron/sqlbuilder