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
	// 返回 404 页面，https://github.com/labstack/echo/issues/671
	echo.NotFoundHandler = func(c echo.Context) error {
		// render your 404 page
		return c.Inline("blog/public/404.html", "404.html")
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
