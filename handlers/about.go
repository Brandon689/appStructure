package handlers

import (
	"github.com/Brandon689/appStructure/views/pages"
	"github.com/labstack/echo/v4"
)

type AboutHandler struct{}

func (h AboutHandler) HandleAbout(c echo.Context) error {

	return render(c, pages.About())
}
