package main

import (
	"blogServer/webpushr"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	tinyUrl "github.com/wingsxdu/tinyurl/example/echo"
)

func main() {
	err := os.Mkdir("./log/", os.ModePerm)
	f, err := os.Create("./log/httpsWarn.log")
	if err != nil {
		panic(err)
	}
	fileWatcher, err := webpushr.NewNotifyFile()
	if err != nil {
		panic(err)
	}
	fileWatcher.WatchDir("./blog/public/post")
	err = webpushr.GetConfig("./config/conf.yaml")
	if err != nil {
		panic(err)
	}
	// 返回 404 页面，https://github.com/labstack/echo/issues/671
	echo.NotFoundHandler = func(c echo.Context) error {
		// render your 404 page
		return c.Redirect(http.StatusTemporaryRedirect, "https://wingsxdu.com/404.html")
	}
	// 在一个协程里监听 HTTP 服务
	go func() {
		f2, err := os.Create("./log/httpWarn.log")
		if err != nil {
			panic(err)
		}
		h := echo.New()
		// 重定向：http://www.wingsxdu.com/ -> https://wingsxdu.com/
		h.Pre(middleware.HTTPSNonWWWRedirect())
		h.Use(middleware.Gzip())
		// 重定向：http://wingsxdu.com/ -> https://wingsxdu.com/
		h.Pre(middleware.HTTPSRedirect())
		// HTTP 服务的日志
		h.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: f2,
		}))
		h.Logger.Warn(h.Start(":80"))
	}()

	e := echo.New()
	// 重定向：https://www.wingsxdu.com/ -> https://wingsxdu.com/
	e.Pre(middleware.NonWWWRedirect())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: f,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: Skipper,
		Root:    "blog/public",
		HTML5:   true,
	}))
	e.HTTPErrorHandler = customHTTPErrorHandler
	tinyUrl.New()
	// 获取 tinyUrl 指向的 url，但是不跳转
	e.GET("/gett/:tinyUrl", tinyUrl.Gett)
	e.GET("/t/:tinyUrl", tinyUrl.GetUrl)
	e.POST("/t", tinyUrl.PostUrl)
	e.PUT("/t", tinyUrl.PutUrl)
	e.DELETE("/t", tinyUrl.DeleteUrl)
	fmt.Printf("当前 PID 为：%d", os.Getpid())
	e.Logger.Warn(e.Start(":443"))
}

type httpError struct {
	code int
	Key  string `json:"error"`
	Msg  string `json:"message"`
}

func customHTTPErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)

	var res = httpError{code: http.StatusInternalServerError, Key: "InternalServerError"}

	if he, ok := err.(*echo.HTTPError); ok {
		res.code = he.Code
		res.Key = http.StatusText(res.code)
		res.Msg = err.Error()
	} else {
		res.Msg = http.StatusText(res.code)
	}

	if !c.Response().Committed {
		err := c.JSON(res.code, res)
		if err != nil {
			c.Logger().Error(err)
		}
	}
}

func Skipper(c echo.Context) bool {
	if strings.HasPrefix(c.Request().RequestURI, "/t") {
		fmt.Println(c.Request().RequestURI)
		return true
	}
	return false
}
