package routes

import (
	"go-first/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine)  {
    //All routes related to users comes here
	router.POST("/auto", controllers.CreateUser())
	router.GET("/auto/:userId", controllers.GetAUser())
	router.PUT("/auto/:userId", controllers.EditAUser())
	router.DELETE("/auto/:userId", controllers.DeleteAUser())
	router.GET("/autos", controllers.GetAllUsers()) 
}