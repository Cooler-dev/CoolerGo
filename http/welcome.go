package http

import (
	"net/http"
	
	"github.com/labstack/echo"
)

func Welcome(c echo.Context) error {
	return c.HTML(http.StatusOK, `
	<!DOCTYPE html>
	<html lang="zh-hans">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>CoolerGo</title>
	</head>
	<body>
		<script>
			console.log("不会吧不会吧，不会真的有人来控制台找彩蛋吧")
		</script>
		<style>
			.main {
				display: flex;
				justify-content: center;
			}
		</style>
		<div class="main">
			<h1>
				<strong>Cooler<span style="color: #06abd7;">Go</span></strong>
			</h1>
		</div>
		<div class="main">
			<p>
				您可以在<a href="https://cooler.js.org">cooler.js.org</a>阅读文档，感谢您的使用！
			</p>
		</div>
	</body>
	</html>
	
	`)
}