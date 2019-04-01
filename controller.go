package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func About(c echo.Context) error {
	return c.String(http.StatusOK, "About")
}
func Contect(c echo.Context) error {
	return c.String(http.StatusOK, "Contect")
}
