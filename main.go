package main

import (
	"log"
	"os"

	"github.com/TaRun1369/goEcommerceCart/controllers"
	"github.com/TaRun1369/goEcommerceCart/database"
	"github.com/TaRun1369/goEcommerceCart/middleware"
	"github.com/TaRun1369/goEcommerceCart/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()      // router ko initialize karna without middleware
	router.Use(gin.Logger()) // middleware ke liye

	routes.UserRoutes(router)
	router.Use(middleware.Authenticate(app))

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
