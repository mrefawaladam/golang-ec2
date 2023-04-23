package route

import (
	"orm/controller"
	m "orm/middleware"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controller.GetUsersController)
	m.LogMiddleware(e)
	e.POST("/users", controller.CreateUser)
	return e
}
