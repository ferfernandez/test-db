package repository

import (
  "database/sql"

  "github.com/ferfernandez/test-db/model"

  _ "github.com/go-sql-driver/mysql"
)

const (
  selectPerson = "SELECT id, firstname, lastname FROM person WHERE id = ?"
  insertPerson = "INSERT INTO person VALUES (?, ?, ?)"
)

type mariaPersonRepository struct {
  factory DBFactory
}

func NewMariaPersonRepository(factory DBFactory) model.PersonRepository  {
  return &mariaPersonRepository{factory: factory}
}

func (m *mariaPersonRepository)GetByID(id uint64) (*model.Person, error) {
  db, err := m.factory.CreateDB()
  if err != nil {
    return nil, err
  }

  defer db.Close()

  var resId int
  var firstname, lastname string
  if err = db.QueryRow(selectPerson, id).Scan(&resId, &firstname, &lastname); err != nil {
    if err == sql.ErrNoRows {
      return nil, nil
    }

    return nil, err
  }

  p := model.Person{
    ID:        id,
    FirstName: firstname,
    LastName:  lastname,
  }

  return &p, nil
}

func (m *mariaPersonRepository)Save(person model.Person) error {
  db, err := m.factory.CreateDB()
  if err != nil {
    return err
  }

  defer db.Close()

  _, err = db.Exec(insertPerson, person.ID, person.FirstName, person.LastName)
  return err
}


