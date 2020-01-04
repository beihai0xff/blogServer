package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	err := os.Mkdir("./log", os.ModePerm)
	f, err := os.Create("./log/httpsWarn.log")
	if err != nil {
		panic(err)
	}
	go func() {
		f2, err := os.Create("./log/httpWarn.log")
		if err != nil {
			panic(err)
		}
		h := echo.New()
		h.Pre(middleware.WWWRedirect())
		h.Pre(middleware.AddTrailingSlash())
		h.Use(middleware.Gzip())
		h.Pre(middleware.HTTPSRedirect())
		h.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: f2,
		}))
		h.Logger.Warn(h.Start(":80"))
	}()

	e := echo.New()
	e.Pre(middleware.WWWRedirect())
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: f,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "blog/public",
		HTML5: true,
	}))
	fmt.Printf("当前 PID 为：%d", os.Getpid())
	e.Logger.Warn(e.StartTLS(":443", "server.pem", "server.key"))
}
