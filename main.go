package main

import (
	// "fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lordscoba/scribbles/middleware"
	routes "github.com/lordscoba/scribbles/routes"
)

func main(){

  // fmt.Println("Hello world")
  err := godotenv.Load(".env")

  if err != nil {
    log.Println("Error loading .env file")
  }
  
  port := os.Getenv("PORT")

  if port==""{
    port="8000"
  }

  router := gin.New()
  router.Use(middleware.CORSMiddleware())
  router.Use(gin.Logger())

  routes.AuthRoutes(router)
  // routes.UserRoutes(router)

  router.GET("/api-1", func(c *gin.Context){
    c.JSON(200, gin.H{"success":"Access granted for api-1"})
  })

  router.GET("/api-2", func(c *gin.Context){
    c.JSON(200, gin.H{"success":"Access granted for api-2"})
  })

  router.Run(":" + port)
}
