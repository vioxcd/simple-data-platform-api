package main

import (
	"vioxcd/dpl/routes"
	"vioxcd/dpl/config"

	"github.com/labstack/echo/v4"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {
	e := echo.New()
	routes.InitRoute(e)
	e.Start(":8000")
}
