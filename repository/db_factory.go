package repository

import (
  "database/sql"
  "fmt"
)

type DBFactory interface {
  // CreateDB creates a DB connection.
  // The one who use this factory, should be responsible for closing
  // the underlying connection.
  CreateDB() (*sql.DB, error)
}

type MariaDBFactory struct {
}

func NewMariaDBFactory() DBFactory {
  return &MariaDBFactory{}
}

func (m *MariaDBFactory)CreateDB() (*sql.DB, error) {
  connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "testuser", "t3stp4ss", "localhost", "test")
  db, err := sql.Open("mysql", connectionString)
  if err != nil {
    return nil, err
  }

  if err = db.Ping(); err != nil {
    db.Close()
    return nil, err
  }

  return db, nil
}
