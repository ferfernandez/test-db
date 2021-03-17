package repository

import (
  "fmt"
  "testing"

  "github.com/DATA-DOG/go-sqlmock"
  "github.com/stretchr/testify/assert"

  "github.com/ferfernandez/test-db/model"
)

const (
  selectTestPattern = "SELECT id, firstname, lastname FROM person WHERE id = \\?"
  insertTestPattern = "INSERT INTO person VALUES \\(\\?, \\?, \\?\\)"
)

func Test_mariaPersonRepository_GetByID_OkFound(t *testing.T) {
  var testID uint64 = 22334455
  testFirstName := "usuario"
  testLastName := "prueba"

  // create & arrange DB Mock
  db, dbMocker := generateMock()

  rows := sqlmock.NewRows([]string{"id", "firstname", "lastname"}).
    AddRow(testID, testFirstName, testLastName)
  dbMocker.ExpectQuery(selectTestPattern).WithArgs(testID).WillReturnRows(rows)

  // create & arrange DBFactory mock,
  // that will return the mocked DB previously created
  mockDBFactory := new(mockDBFactory)
  mockDBFactory.On("CreateDB").Return(db, nil)

  // effectively do the test...
  repo := &mariaPersonRepository{factory: mockDBFactory}
  expectedResult := model.Person{
    ID:        testID,
    FirstName: testFirstName,
    LastName:  testLastName,
  }

  result, err := repo.GetByID(testID)

  // check result
  assert.NoError(t, err)
  assert.Equal(t, expectedResult, *result)
}

func Test_mariaPersonRepository_GetByID_OkNotFound(t *testing.T) {
  var testID uint64 = 22334455

  // create & arrange DB Mock
  db, dbMocker := generateMock()

  rows := sqlmock.NewRows([]string{"id", "firstname", "lastname"})
  dbMocker.ExpectQuery(selectTestPattern).WithArgs(testID).WillReturnRows(rows)

  // create & arrange DBFactory mock,
  // that will return the mocked DB previously created
  mockDBFactory := new(mockDBFactory)
  mockDBFactory.On("CreateDB").Return(db, nil)

  // effectively do the test...
  repo := &mariaPersonRepository{factory: mockDBFactory}

  result, err := repo.GetByID(testID)

  // check result
  assert.NoError(t, err)
  assert.Nil(t, result)
}

func Test_mariaPersonRepository_GetByID_FailCreateDB(t *testing.T) {
  var testID uint64 = 22334455
  expectedError := fmt.Errorf("some error while creating db")

  // create & arrange DBFactory mock
  mockDBFactory := new(mockDBFactory)
  mockDBFactory.On("CreateDB").Return(nil, expectedError)

  // effectively do the test...
  repo := &mariaPersonRepository{factory: mockDBFactory}

  _, err := repo.GetByID(testID)
  // check result
  assert.Error(t, err)
  assert.Equal(t, expectedError, err)
}

func Test_mariaPersonRepository_GetByID_FailQuery(t *testing.T) {
  var testID uint64 = 22334455
  expectedError := fmt.Errorf("some error while executing query")

  // create & arrange DB Mock
  db, dbMocker := generateMock()
  dbMocker.ExpectQuery(selectTestPattern).WithArgs(testID).WillReturnError(expectedError)

  // create & arrange DBFactory mock,
  // that will return the mocked DB previously created
  mockDBFactory := new(mockDBFactory)
  mockDBFactory.On("CreateDB").Return(db, nil)

  // effectively do the test...
  repo := &mariaPersonRepository{factory: mockDBFactory}

  _, err := repo.GetByID(testID)

  // check result
  assert.Error(t, err)
  assert.Equal(t, expectedError, err)
}

func Test_mariaPersonRepository_Save_Ok(t *testing.T) {
  var testID uint64 = 22334455
  testFirstName := "usuario"
  testLastName := "prueba"

  // create & arrange DB Mock
  db, dbMocker := generateMock()
  dbMocker.ExpectExec(insertTestPattern).WithArgs(testID, testFirstName, testLastName).WillReturnResult(sqlmock.NewResult(0,1))

  // create & arrange DBFactory mock,
  // that will return the mocked DB previously created
  mockDBFactory := new(mockDBFactory)
  mockDBFactory.On("CreateDB").Return(db, nil)

  // effectively do the test...
  repo := &mariaPersonRepository{factory: mockDBFactory}
  p := model.Person{
    ID:        testID,
    FirstName: testFirstName,
    LastName:  testLastName,
  }

  err := repo.Save(p)

  // check result
  assert.NoError(t, err)
}

func Test_mariaPersonRepository_Save_FailExec(t *testing.T) {
  var testID uint64 = 22334455
  testFirstName := "usuario"
  testLastName := "prueba"

  expectedError := fmt.Errorf("some error while exec")

  // create & arrange DB Mock
  db, dbMocker := generateMock()
  dbMocker.ExpectExec(insertTestPattern).WithArgs(testID, testFirstName, testLastName).WillReturnError(expectedError)

  // create & arrange DBFactory mock,
  // that will return the mocked DB previously created
  mockDBFactory := new(mockDBFactory)
  mockDBFactory.On("CreateDB").Return(db, nil)

  // effectively do the test...
  repo := &mariaPersonRepository{factory: mockDBFactory}
  p := model.Person{
    ID:        testID,
    FirstName: testFirstName,
    LastName:  testLastName,
  }

  err := repo.Save(p)

  // check result
  assert.Error(t, err)
  assert.Equal(t, expectedError, err)
}

func Test_mariaPersonRepository_Save_FailCreateDB(t *testing.T) {
  var testID uint64 = 22334455
  testFirstName := "usuario"
  testLastName := "prueba"

  expectedError := fmt.Errorf("some error while creating db")

  // create & arrange DBFactory mock,
  // that will return the mocked DB previously created
  mockDBFactory := new(mockDBFactory)
  mockDBFactory.On("CreateDB").Return(nil, expectedError)

  // effectively do the test...
  repo := &mariaPersonRepository{factory: mockDBFactory}
  p := model.Person{
    ID:        testID,
    FirstName: testFirstName,
    LastName:  testLastName,
  }

  err := repo.Save(p)

  // check result
  assert.Error(t, err)
  assert.Equal(t, expectedError, err)
}

