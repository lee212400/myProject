package router

import (
	"myProject/src/interface/handler"

	"github.com/labstack/echo"
)

func InitOrderRouting(e *echo.Echo, orderHandler *handler.OrderHandler) {
	e.POST("/order", orderHandler.Post())
	e.GET("/order/:id", orderHandler.Get())
}
