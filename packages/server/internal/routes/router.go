package routes

import "github.com/labstack/echo/v5"

func RegisterRoutes(e *echo.Echo) {
	e.GET("/health", HealthCheck)
}

func HealthCheck(c *echo.Context) error {
	return c.String(200, "OK")
}
