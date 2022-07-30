package http

import (
	"net/http"
	"os"

	"github.com/Cooler-dev/CoolerGo/lib"
	"github.com/labstack/echo"
)

func Welcome(c echo.Context) error {
	file, err := os.Open("D:/gopath/src/golang_development_notes/example/log.txt")
	if err != nil {
		log.Error("文件读取错误")
	}
	defer file.Close()
}