package routes

import (
	"golang-api/controllers"

	"github.com/labstack/echo"
)

func Route() {
	e := echo.New()

	e.GET("/", controllers.Home)
	e.GET("/about", controllers.About)
	e.GET("/about", controllers.About)

	e.Logger.Fatal(e.Start(":8080"))
}
