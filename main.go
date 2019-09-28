package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	f, err := os.Create("./log/test.log")
	if err != nil {
		panic(err)
	}
	//logFile, _ := os.OpenFile("./log/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: f,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Pre(middleware.WWWRedirect())
	e.Pre(middleware.AddTrailingSlash())
	e.Pre(middleware.HTTPSRedirect())

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "blog/public",
		HTML5: true,
	}))
	e.Logger.Fatal(e.StartTLS(":443", "server.crt", "server.key"))
}
