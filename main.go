package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wildanfaz/websocket_app/producer"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Static("/", "../public")
	e.GET("/ws", producer.Write)
	e.Logger.Fatal(e.Start(":3030"))
}
