package cmd

import (
	"fmt"

	"github.com/Cooler-dev/CoolerGo/lib"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() {
	fmt.Println(`  ____           _            ____      `)
	fmt.Println(`/ ___| ___   ___ | | ___ _ __ / ___|  ___  `)
	fmt.Println(`| |   / _ \ / _ \| |/ _ \ '__| |  _ / _ \ `)
	fmt.Println(`| |__| (_) | (_) | |  __/ |  | |_| | (_) |`)
	fmt.Println(` \____\___/ \___/|_|\___|_|   \____|\___/ `)
	fmt.Println(`===========================================`)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	lib.Route()
	e.Logger.Fatal(e.Start("localhost:3000"))
}