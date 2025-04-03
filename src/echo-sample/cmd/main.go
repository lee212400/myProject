package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	newE := echo.New()
	newE.GET("/user:name", func(e echo.Context) error {
		name := e.Param("name")
		return e.String(http.StatusOK, fmt.Sprintf("Hello, %s User!", &name))
	})

	newE.POST("/user", func(e echo.Context) error {
		var reqDt any
		if err := e.Bind(&reqDt); err != nil {
			return e.JSON(http.StatusBadRequest, err.Error())
		}

		return e.String(http.StatusOK, "User Submitted!")
	})

	newE.Logger.Fatal(newE.Start(":8888"))
}
