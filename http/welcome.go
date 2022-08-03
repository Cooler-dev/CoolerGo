package http

import (
	"net/http"
	
	"github.com/labstack/echo"
)

func Welcome(c echo.Context) error {
	data := map[string]string{"message":"欢迎使用CoolerGo"}
	return c.JSON(http.StatusOK,data)
}