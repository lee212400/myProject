package main

import (
	"fmt"
	"net/http"
	"regexp"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func main() {
	newE := echo.New()

	//routing(newE)
	//request(newE)
	validateSample(newE)

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
	UserName  string `json:"userName" validate:"required"`
	UserAge   int    `json:"userAge" validate:"required"`
	UserEmail string `json:"userEmail" validate:"required,email"`
	UserId    string `json:"userId" validate:"required,email"`
}

type CustomValidator struct {
	validator *validator.Validate
}

var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9]{1,20}$`)

func (cv *CustomValidator) Validate(i interface{}) error {
	// http status制御
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func validateSample(e *echo.Echo) {
	validate := validator.New()
	// keyを追加して、特定項目に対してカスタマイズ可能
	validate.RegisterValidation("userName", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		return len(val) >= 3 && len(val) <= 20
	})

	// カスタマイズで正規表現のvalidateチェック可能
	validate.RegisterValidation("userId", func(fl validator.FieldLevel) bool {
		return usernameRegex.MatchString(fl.Field().String())
	})

	// 単一validateチェック
	e.GET("/user/:name", func(e echo.Context) error {
		name := e.Param("name")
		if err := validate.Var(name, "required,min=2,max=10"); err != nil {
			return e.String(http.StatusBadRequest, err.Error())
		}
		return e.String(http.StatusOK, fmt.Sprintf("Hello, %s User!", name))
	})

	// 構造体を利用してvalidateチェック
	e.PUT("user/:id", func(e echo.Context) error {
		id := e.Param("id")
		reqDt := UserRequest{}

		if err := e.Bind(&reqDt); err != nil {
			return e.JSON(http.StatusBadRequest, err.Error())
		}

		if err := validate.Struct(reqDt); err != nil {
			return e.String(http.StatusBadRequest, err.Error())
		}

		res := fmt.Sprintf("id:%s, name:%s, age:%d", id, reqDt.UserName, reqDt.UserAge)
		return e.String(http.StatusOK, res)
	})

	// validate customer
	e.Validator = &CustomValidator{validator: validate}
	e.PUT("user/:id", func(c echo.Context) error {
		id := c.Param("id")
		reqDt := UserRequest{}

		if err := c.Bind(&reqDt); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(reqDt); err != nil {
			return err
		}

		res := fmt.Sprintf("id:%s, name:%s, age:%d", id, reqDt.UserName, reqDt.UserAge)
		return c.String(http.StatusOK, res)
	})

}
