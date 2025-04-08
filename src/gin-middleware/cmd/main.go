package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	setMiddleware(router)
	router.Run(":8080")
}

func setMiddleware(g *gin.Engine) {
	// revocer(panic対応)
	g.Use(gin.Recovery())

	// CORS
	g.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://your-frontend.com"},
		AllowMethods: []string{"GET", "POST"},
	}))

	// loggingカスタマイズ
	g.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		fmt.Printf("%s %s %v\n", c.Request.Method, c.Request.URL.Path, latency)
	})

	// 個別ミドルウェア設定
	g.POST("/upload",
		RateLimitMiddleware(100, time.Minute), // rate limit
		FileSizeLimitMiddleware(10<<20),       // file size limit
		func(c *gin.Context) {
			file, _ := c.FormFile("file")
			c.JSON(200, gin.H{"filename": file.Filename})
		},
	)
}

var (
	requests = make(map[string]int)
	mu       sync.Mutex
)

func RateLimitMiddleware(limit int, duration time.Duration) gin.HandlerFunc {

	return func(c *gin.Context) {
		// limit処理追加

		c.Next()
	}
}

func FileSizeLimitMiddleware(maxBytes int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)

		if err := c.Request.ParseMultipartForm(maxBytes); err != nil {
			c.AbortWithStatusJSON(http.StatusRequestEntityTooLarge, gin.H{
				"error": "file size limit)",
			})
			return
		}

		c.Next()
	}
}
