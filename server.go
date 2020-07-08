package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// サーバー起動
	e.Start(":8080")
}