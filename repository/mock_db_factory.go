package repository

import (
  "database/sql"
  "log"

  "github.com/DATA-DOG/go-sqlmock"
  "github.com/stretchr/testify/mock"
)

type mockDBFactory struct {
  mock.Mock
}

func (k *mockDBFactory)CreateDB() (*sql.DB, error) {
  args := k.Called()

  if args.Get(0) == nil {
    return nil, args.Error(1)
  }

  return args.Get(0).(*sql.DB), args.Error(1)
}

// generateMock is an utilitarian function for generating fake DB an mock objects
func generateMock() (*sql.DB, sqlmock.Sqlmock) {
  db, mock, err := sqlmock.New()
  if err != nil {
    log.Fatalf("couldn't generate mock object")
  }

  return db, mock
}
