package api

import (
	"github.com/Brandon689/appStructure/types"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetJSONData(c echo.Context) error {
	user := types.User{
		ID:   1,
		Name: "John Doe",
	}
	return c.JSON(http.StatusOK, user)
}
