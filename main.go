package main

import (
	"fmt"
	"go-qr-app/model"

	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

//DB接続を想定してTODO固定値で実装

var todos = []model.Todo{
	{
		ID:        rand.Int(),
		Text:      "test1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{ID: rand.Int(), Text: "test2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: rand.Int(), Text: "test3", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: rand.Int(), Text: "test4", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"todos": todos})
	})

	router.GET("/detail/:id", func(ctx *gin.Context) {
		p := ctx.Param("id")
		id, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}

		var todo = GetOne(id)
		fmt.Println("todo出します")
		fmt.Println(todo)

		ctx.HTML(200, "detail.html", gin.H{"todo": todo})
	})

	// router.POST("/new", func(ctx *gin.Context) {
	// 	ID := ctx.PostForm("ID")
	// 	text := ctx.PostForm("Text")
	// 	var newTodo
	// 	ctx.Redirect(302, "/")
	// })

	// 	router.POST("/update/:id", func(ctx *gin.Context) {
	// 		n := ctx.Param("id")
	// 		id, err := strconv.Atoi(n)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		text := ctx.PostForm("text")
	// 		status := ctx.PostForm("status")
	// 		dbUpdate(id, text, status)
	// 		ctx.Redirect(302, "/")
	// 	})

	// })

	//削除確認
	// router.GET("/delete_check/:id", func(ctx *gin.Context) {
	// 	n := ctx.Param("id")
	// 	id, err := strconv.Atoi(n)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	todo := dbGetOne(id)
	// 	ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	// })

	// router.POST("/delete/:id", func(ctx *gin.Context) {
	// 	n := ctx.Param("id")
	// 	id, err := strconv.Atoi(n)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	dbDeleta(id)
	// 	ctx.Redirect(302, "/")
	// })
	router.Run()
}

func GetOne(targetID int) *model.Todo {
	for _, todo := range todos {
		if todo.ID == targetID {
			return &todo
		}
	}
	return nil
}
