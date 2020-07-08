package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// サーバー起動
	e.Start(":8080")
	fmt.Println("server start")
}