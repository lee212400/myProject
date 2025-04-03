package main

import (
	"github.com/lee212400/myProject/infrastructure/repository"
	"github.com/lee212400/myProject/interface/handler"
	"github.com/lee212400/myProject/interface/router"
	"github.com/lee212400/myProject/usecase"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// order
	orderRepository := repository.NewOrderRepository()
	orderInteractor := usecase.NewOrderInteractor(orderRepository)
	orderHandler := handler.NewOrderHandler(orderInteractor)
	router.InitOrderRouting(e, orderHandler)

	e.Logger.Fatal(e.Start(":8888"))
}
