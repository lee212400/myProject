package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	newE := echo.New()

	//routing(newE)
	request(newE)

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

func request(e *echo.Echo) {
	// 経路
	e.GET("/user/:name", func(e echo.Context) error {
		name := e.Param("name")
		return e.String(http.StatusOK, fmt.Sprintf("Hello, %s User!", name))
	})

	// 経路 + from-data
	e.PUT("user/:id", func(e echo.Context) error {
		id := e.Param("id")
		name := e.FormValue("name")
		age := e.FormValue("age")

		res := fmt.Sprintf("id:%s, name:%s, age:%s", id, name, age)
		return e.String(http.StatusOK, res)
	})

	// query
	e.GET("/user/search", func(e echo.Context) error {
		name := e.QueryParam("name")
		limit := e.Param("limit")

		res := fmt.Sprintf("name:%s, limit:%s", name, limit)
		return e.String(http.StatusOK, res)
	})

	// json形式でbinding
	e.POST("/user", func(e echo.Context) error {
		var reqDt UserRequest
		if err := e.Bind(&reqDt); err != nil {
			return e.JSON(http.StatusBadRequest, err.Error())
		}
		res := fmt.Sprintf("name:%s, age:%d", reqDt.UserName, reqDt.UserAge)
		return e.String(http.StatusOK, res)
	})

	// bindingを構造体を使わないでできるか確認
	e.POST("/user", func(e echo.Context) error {
		var reqDt any
		if err := e.Bind(&reqDt); err != nil {
			return e.JSON(http.StatusBadRequest, err.Error())
		}

		// ひとつづつ取り出したり、項目確認するが面倒くさい
		// 構造体で管理した方が良さそう
		name := ""
		age := 0
		dt := reqDt.(map[string]any)
		if v, ok := dt["userName"].(string); ok {
			name = v
		}
		if v, ok := dt["userAge"].(int); ok {
			age = v
		}
		res := fmt.Sprintf("name:%s, age:%d", name, age)

		return e.String(http.StatusOK, res)
	})
}

type UserRequest struct {
	UserName string `json:"userName"`
	UserAge  int    `json:"userAge"`
}
