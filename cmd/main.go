package main

import (
	"github.com/Brandon689/appStructure/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Static("/", "assets")

	aboutHandler := &handlers.AboutHandler{}
	homeHandler := &handlers.HomeHandler{}
	handlers.SetupRoutes(e, homeHandler, aboutHandler)

	e.Use(middleware.Logger())

	// Start Server
	e.Logger.Fatal(e.Start(":8082"))
}
