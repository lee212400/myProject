package main

import (
	"myProject/src/DDD-sample/application"
	"myProject/src/DDD-sample/infrastructure/repository"
	"myProject/src/DDD-sample/interface/handler"
	"myProject/src/DDD-sample/interface/router"

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
