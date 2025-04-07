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
	// request(router)
	requestStruct(router)

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

type User struct {
	Name  string `json:"name" form:"name" uri:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required,email"`
}

type Header struct {
	Token string `header:"token" binding:"required"`
}

func requestStruct(g *gin.Engine) {
	// json
	g.POST("/user", func(c *gin.Context) {
		var u User
		if err := c.ShouldBindJSON(&u); err != nil {
			c.String(http.StatusBadRequest, "error: %v", err.Error())
			return
		}
		c.String(http.StatusOK, "OK")
	})

	// query parameter
	g.GET("/user", func(c *gin.Context) {
		var u User
		if err := c.ShouldBindQuery(&u); err != nil {
			c.String(http.StatusBadRequest, "error: %v", err.Error())
			return
		}
		c.String(http.StatusOK, "OK")
	})

	// uri経路
	g.GET("/user/:name/:email", func(c *gin.Context) {
		var u User
		if err := c.ShouldBindUri(&u); err != nil {
			c.String(http.StatusBadRequest, "error: %v", err.Error())
			return
		}
		c.String(http.StatusOK, "OK")
	})

	// content-typeによってbinding
	// jsonの場合josn、application/x-www-form-urlencoded場合form
	g.POST("/user", func(c *gin.Context) {
		var u User
		if err := c.ShouldBind(&u); err != nil {
			c.String(http.StatusBadRequest, "error: %v", err.Error())
			return
		}
		c.String(http.StatusOK, "OK")
	})

	// headerのbinding
	g.GET("/header", func(c *gin.Context) {
		var h Header
		if err := c.ShouldBindHeader(&h); err != nil {
			c.String(http.StatusBadRequest, "error: %v", err.Error())
			return
		}
		c.String(http.StatusOK, "OK")
	})
}
