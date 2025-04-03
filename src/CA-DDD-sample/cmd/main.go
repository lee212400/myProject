package main

import (
	"github.com/lee212400/myProject/application"
	"github.com/lee212400/myProject/infrastructure/repository"
	"github.com/lee212400/myProject/interface/handler"
	"github.com/lee212400/myProject/interface/router"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// order
	orderRepository := repository.NewOrderRepository()
	orderService := application.NewOrderService(orderRepository)
	orderHandler := handler.NewOrderHandler(orderService)
	router.InitOrderRouting(e, orderHandler)

	e.Logger.Fatal(e.Start(":8888"))
}
