package main

import (
	"context"
	"fmt"
	"github.com/Brandon689/appStructure/db"
	"github.com/Brandon689/appStructure/handlers"
	"github.com/Brandon689/appStructure/types"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	//util.GetJsonExternalAPIExample()

	ctx := context.Background()

	db2, err := db.NewDatabase()
	if err != nil {
		panic(err)
	}

	err = db2.Migrate(ctx)
	if err != nil {
		panic(err)
	}

	user := &types.User{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}
	err = db2.InsertUser(ctx, user)
	if err != nil {
		panic(err)
	}

	user, err = db2.GetUserByName(ctx, "John Doe")
	if err != nil {
		panic(err)
	}

	fmt.Printf("User ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)

	e := echo.New()

	e.Static("/", "assets")

	aboutHandler := &handlers.AboutHandler{}
	homeHandler := &handlers.HomeHandler{}
	handlers.SetupRoutes(e, homeHandler, aboutHandler)

	e.Use(middleware.Logger())

	// Start Server
	e.Logger.Fatal(e.Start(":8082"))
}
