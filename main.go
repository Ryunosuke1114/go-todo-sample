package main

import (
	"fmt"
	"go-qr-app/handler"
	"go-qr-app/infra"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		var todos = handler.NewTodo(infra.NewTodoRepo()).Get()
		ctx.HTML(200, "index.html", gin.H{"todos": todos})
	})

	router.GET("/detail/:id", func(ctx *gin.Context) {
		p := ctx.Param("id")
		targetID, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		var todo = handler.NewTodo(infra.NewTodoRepo()).GetOne(targetID)
		ctx.HTML(200, "detail.html", gin.H{"todo": todo})
	})

	router.POST("/new", func(ctx *gin.Context) {
		id := rand.Int()
		text := ctx.PostForm("text")
		handler.NewTodo(infra.NewTodoRepo()).Create(id, text)
		todos := handler.NewTodo(infra.NewTodoRepo()).Create(id, text)
		ctx.HTML(200, "index.html", gin.H{"todos": todos})
	})

	router.POST("/update/:id", func(ctx *gin.Context) {
		fmt.Println("PUT now")
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		text := ctx.PostForm("text")
		handler.NewTodo(infra.NewTodoRepo()).Update(id, text)
		ctx.Redirect(302, "/")
	})

	// 	router.POST("/delete/:id", func(ctx *gin.Context) {
	// 		n := ctx.Param("id")
	// 		id, err := strconv.Atoi(n)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		DeleteTodo(id)
	// 		ctx.Redirect(302, "/")
	// 	})
	router.Run()
}

