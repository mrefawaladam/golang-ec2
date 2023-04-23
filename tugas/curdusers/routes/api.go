package routes

import (
	"curdusers/constans"
	"curdusers/controllers"
	m "curdusers/middlewares"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	// Auth
	e.POST("/login", controllers.LoginUserController)

	// blogs routes
	e.POST("/blogs", controllers.CreateBlogController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)

	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(mid.BasicAuth(m.BasicAuthDB))
	eAuthBasic.GET("/users", controllers.GetUsersController)

	eJwt := e.Group("/jwt")
	eJwt.Use(mid.JWT([]byte(constans.SECRET_JWT)))
	eJwt.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	eJwt.POST("/users", controllers.CreateUserController)
	eJwt.DELETE("/users/:id", controllers.DeleteUserController)
	eJwt.PUT("/users/:id", controllers.UpdateUserController)

	// Book routes
	eJwt.POST("/books", controllers.CreateBookController)
	eJwt.GET("/books", controllers.GetBooksController)
	eJwt.GET("/books/:id", controllers.GetBookController)
	eJwt.PUT("/books/:id", controllers.UpdateBookController)
	eJwt.DELETE("/books/:id", controllers.DeleteBookController)

	return e
}
