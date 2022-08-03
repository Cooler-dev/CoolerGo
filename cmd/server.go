package cmd

import (
	"fmt"

	"github.com/labstack/echo"
)

func Start() {
	banner := `
 _______                    _               
/ ______|                  | |              
| |        _____   _____   | |  ___   _____
| |       /  _  \ /  _  \  | | /  _ \ | '__|
| |_____  | (_) | | (_) |  | | |  __/ | |   
\_______| \_____/ \_____/  |_| \____  |_|  

=============================================
	`

	e := echo.New()
	fmt.Println(banner)
	InitRoute(e)
	InitLog()
	e.Logger.Fatal(e.Start("localhost:3000"))
}