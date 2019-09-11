package route

import (
	"github.com/labstack/echo/v4"
	"leiax00.com/fxWeb/handler"
)

func RegisterRoutes(handler *handler.WebHandler, e *echo.Echo) {
	group := e.Group("/web")
	group.GET("/hello", handler.HelloApp)
	e.Logger.Info("Success to register routes!!!")
}