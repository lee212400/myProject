package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	newE := echo.New()

	routing(newE)

	newE.Logger.Fatal(newE.Start(":8888"))
}

func routing(e *echo.Echo) {
	userIdGroup := e.Group("/user") // 重複するprefixはgroup化した方が管理しやすい

	// 経路とパラメータ設定
	e.GET("/user/:name", func(e echo.Context) error {
		name := e.Param("name")
		return e.String(http.StatusOK, fmt.Sprintf("Hello, %s User!", name))
	})

	// 経路設定
	e.POST("/user", func(e echo.Context) error {
		var reqDt any
		if err := e.Bind(&reqDt); err != nil {
			return e.JSON(http.StatusBadRequest, err.Error())
		}

		return e.String(http.StatusOK, "User Submitted!")
	})

	// 経路とパラメータ設定
	userIdGroup.PUT("/:id", func(e echo.Context) error {
		var reqDt any
		id := e.Param("id")
		if err := e.Bind(&reqDt); err != nil {
			return e.JSON(http.StatusBadRequest, err.Error())
		}

		return e.String(http.StatusOK, fmt.Sprintf("PUT, %s User!", id))
	})

	// 経路とパラメータ設定
	userIdGroup.DELETE("/:id", func(e echo.Context) error {
		id := e.Param("id")

		return e.String(http.StatusOK, fmt.Sprintf("DELETE, %s User!", id))
	})

}
