package handler

import (
  "net/http"
  "strconv"

  "github.com/ferfernandez/test-db/model"
  "github.com/ferfernandez/test-db/repository"
  "github.com/gin-gonic/gin"
)

var personRepository model.PersonRepository


func init() {
  personRepository = repository.NewMariaPersonRepository(repository.NewMariaDBFactory())
}

func GetPerson(c *gin.Context) {
  id, e := strconv.ParseUint(c.Param("id"), 10, 64)

  if e != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
    return
  }

  p, e := personRepository.GetByID(id)

  if e != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
    return
  }

  if p == nil {
    c.AbortWithStatus(http.StatusNotFound)
    return
  }

  c.JSON(http.StatusOK, p)
}

func PostPerson(c *gin.Context) {
  var person model.Person

  if err := c.ShouldBindJSON(&person); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  if err := personRepository.Save(person); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusCreated, person)
}
