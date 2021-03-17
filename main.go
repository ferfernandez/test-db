package main

import (
  "github.com/gin-gonic/gin"
  "github.com/ferfernandez/test-db/handler"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

  g := r.Group("/person")
  {
    g.GET("/:id", handler.GetPerson)
    g.POST("/", handler.PostPerson)
  }

  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
