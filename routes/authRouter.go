package routes

import(
	controller "github.com/lordscoba/scribbles/controllers"
	"github.com/gin-gonic/gin"
	middleware "github.com/lordscoba/scribbles/middleware"
)

func AuthRoutes(incomingRoutes *gin.Engine){

	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
	incomingRoutes.POST("/", controller.Page())
}
