package main

import (
	"github.com/lee212400/myProject/interface/router"
	"github.com/lee212400/myProject/wire"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// order
	orderHandler := wire.InitOrder()
	router.InitOrderRouting(e, orderHandler)

	e.Logger.Fatal(e.Start(":8888"))
}
