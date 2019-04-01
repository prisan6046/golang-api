package main

import (
	"github.com/labstack/echo"
)

func Route() {
	e := echo.New()

	e.GET("/", Home)
	e.GET("/about", About)
	e.Logger.Fatal(e.Start(":8080"))
}
