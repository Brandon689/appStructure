package handlers

import "github.com/labstack/echo/v4"

func SetupRoutes(e *echo.Echo, hh *HomeHandler, ah *AboutHandler) {
	e.GET("/", hh.HandleHome)
	e.GET("/about", ah.HandleAbout)
}
