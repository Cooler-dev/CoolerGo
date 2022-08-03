package cmd

import (
	"github.com/Cooler-dev/CoolerGo/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRoute(e *echo.Echo) {
	e.GET("/", http.Welcome)
}

func InitLog() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}