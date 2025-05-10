package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	//setMiddleware(e)
	testMiddlewar(e)
	e.Logger.Fatal(e.Start(":8888"))
}
func setMiddleware(e *echo.Echo) {
	// rogging + custom
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${method} ${uri} ${status} ${latency_human}\n`,
	}))

	e.Use(middleware.Recover()) // panic発生したら、サーバーが落ちないようにする

	// CORS許可
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://sample.com"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	// jwt token認証
	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte("my-secret"),
		ContextKey:  "user",
		TokenLookup: "header:Authorization", // "Authorization: Bearer <token>"
		KeyFunc:     getKey,
	})

	e.GET("/user/:id", func(c echo.Context) error {
		myId := c.Param("id")
		// JWTから情報抽出
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(string)
		if myId != id {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Unauthorized",
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": "OK id" + id,
		})
	}, jwtMiddleware)

	// api limit設定
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	// header request Id挿入
	e.Use(middleware.RequestID())
}

// keyfuncをカスタマイズしてjwtのヘッダー,exipre等確認ができる
func getKey(token *jwt.Token) (interface{}, error) {
	_, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have a key ID in the kid field")
	}

	typ, ok := token.Header["typ"].(string)
	if ok {
		fmt.Printf("Token type (typ): %s\n", typ)
	} else {
		fmt.Println("No 'typ' field in the JWT header")
	}

	return nil, nil
}

func myMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// ミドルウェア: requestされたらログ出力
		fmt.Println("test myMiddleware")
		return next(c)
	}
}

func myGroupMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// ミドルウェア: requestされたらログ出力
		fmt.Println("test myGroupMiddleware")
		return next(c)
	}
}

func testMiddlewar(e *echo.Echo) {
	// router個別にミドルウェア設定
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is a secured API route")
	}, myMiddleware)

	// ミドルウェア設定なし
	e.GET("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is a secured API route")
	})

	// routerグループことにミドルウェア設定
	apiGroup := e.Group("/api", myGroupMiddleware)

	// /api/hello rout
	apiGroup.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from API group")
	})

	// /api/world rout
	apiGroup.GET("/world", func(c echo.Context) error {
		return c.String(http.StatusOK, "World from API group")
	})
}
