package cmd

import (
	"fmt"

	"github.com/Cooler-dev/CoolerGo/lib"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() {
	fmt.Println("欢迎使用CoolerGo!")
	fmt.Println("Welcome to use CoolerGo!")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	lib.Route()
	e.Logger.Fatal(e.Start(":3000"))
}