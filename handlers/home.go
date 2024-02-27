package handlers

import (
	"github.com/Brandon689/appStructure/views/pages"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (h HomeHandler) HandleHome(c echo.Context) error {

	return render(c, pages.Home())
}
