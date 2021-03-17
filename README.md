# Sample Golang app for DB access

> This app shows how to access to DB, using the go built-in package `database/sql`. This source code was developed with Intellij IDEA + Golang plugin.

## Specifications
These are the externals tools (external to golang) that were used for this app:
* [gin-gonic](https://github.com/gin-gonic/gin)
* [MySQL Go Driver](https://github.com/go-sql-driver/mysql)
* [Sql driver Mock for Golang](https://github.com/DATA-DOG/go-sqlmock)
* A [MariaDB docker instance](https://hub.docker.com/_/mariadb). You can do it by your own following a couple of tutorials over internet, like [this one](https://mariadb.com/kb/en/installing-and-using-mariadb-via-docker/). Sorry for not _dockerizing_ the app, and composing both (app + MariaDB) with docker compose.

## How to build It? Run It? Run tests?
You can do that via IDE, using Intellij IDEA + Golang plugin, or Goland IDE.

Or, you can do it via command-line. 

### Build

```
> cd test-db
> go build ./...
```

### Run
Run main file directly:

```
> go run main.go
```

### Run tests

```
> go test -v ./...
```

