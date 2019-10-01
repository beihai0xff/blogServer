package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"os"
)

func main() {
	err := os.Mkdir("./log", os.ModePerm)
	go func() {
		f2, err := os.Create("./log/httpWarn.log")
		if err != nil {
			panic(err)
		}
		h := echo.New()
		h.Pre(middleware.WWWRedirect())
		h.Pre(middleware.AddTrailingSlash())
		h.Pre(middleware.HTTPSRedirect())
		h.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: f2,
		}))
		h.Logger.Warn(h.Start(":80"))
	}()

	f, err := os.Create("./log/httpsWarn.log")
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Pre(middleware.WWWRedirect())
	e.Pre(middleware.AddTrailingSlash())
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
	e.Logger.SetLevel(log.WARN)
	e.Logger.Warn(e.StartTLS(":443", "server.crt", "server.key"))
}
