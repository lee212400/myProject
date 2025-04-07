package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//	routing(router)
	request(router)

	router.Run(":8888")
}

func routing(g *gin.Engine) {
	g.Group("/user")
	g.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Name: %s", name)
	})

	g.POST("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Name: %s", name)
	})

	g.PUT("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Name: %s", name)
	})

	g.DELETE("/user/:id", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Name: %s", name)
	})

	g.PATCH("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Name: %s", name)
	})

	g.HEAD("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Name: %s", name)
	})

	g.OPTIONS("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Name: %s", name)
	})
}

func request(g *gin.Engine) {
	g.Use(gin.Recovery()) // paninc recover

	// パラメータ取得
	g.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		if name == "" {
			c.String(http.StatusBadRequest, "name required")
			return
		}
		c.String(http.StatusOK, "Hello %s", name)
	})

	// Query取得
	// /user?firstname=test&lastname=name
	g.GET("/user", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest") // firstNameのデフォルト値設定
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "firstName: %s, lastName: %s", firstname, lastname)
	})

	// Multipart/Urlencoded Form
	g.POST("/user", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// Map as querystring or postform parameters
	g.POST("/user", func(c *gin.Context) {
		// /user?ids[a]=12&ids[b]=123
		// body names[first]=test1&names[second]=test2
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.String(http.StatusOK, "ids: %s, names: %s", ids, names)
	})

	g.POST("/file", func(c *gin.Context) {
		// Single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "dst")

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
}
