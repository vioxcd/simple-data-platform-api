package routes

import (
	"os"
	"vioxcd/dpl/controllers"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())

	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	e.POST("/user", controllers.AddUser)
	e.POST("/login", controllers.Login)

	eAuth.GET("/log", controllers.GetRunsHistory)
	eAuth.POST("/log", controllers.RunNewSnapshot)

	eAuth.GET("/snapshot", controllers.GetSnapshots)
}
