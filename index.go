package main

import (
	"go-first/configs"
	"go-first/routes"

	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  //run database
  configs.ConnectDB()

  //routes
  routes.UserRoute(router)

  router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "data": "Hello from Gin-gonic & mongoDB",
    })
  })

  router.Run("localhost:3461") 
}

// mongodb://localhost:27017/?readPreference=primary&directConnection=true&ssl=false