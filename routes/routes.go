package routes

import (
	"github.com/TaRun1369/goEcommerceCart/controllers"
	"github.com/gin-goric/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.Post("/admin/addproduct", controllers.GetUser())
	incomingRoutes.GET("/users/productview", controllers.GetAllUsers())
	incomingRoutes.GET("/users/search", controllers.UpdateUser())
}