package model

type (
  Person struct {
    ID        uint64  `json:"id"`
    FirstName string  `json:"first_name"`
    LastName  string  `json:"last_name"`
  }

  PersonRepository interface {
    GetByID(id uint64) (*Person, error)
    Save(person Person) error
  }
)

