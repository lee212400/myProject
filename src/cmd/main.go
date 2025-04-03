package main

import (
	"myProject/src/application"
	"myProject/src/infrastructure/repository"
	"myProject/src/interface/handler"
	"myProject/src/interface/router"

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
