package lib

import (
	"github.com/labstack/echo"
	"github.com/Cooler-dev/CoolerGo/http"
)

func Route() {
	e := echo.New()
	e.GET("/", http.Welcome())
}