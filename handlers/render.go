package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// render creates a compatible interface for rendering templ components
// using the Echo web framework.
func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
