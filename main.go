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
		fmt.Println(id)
		text := ctx.PostForm("text")
		fmt.Println(text)
		handler.NewTodo(infra.NewTodoRepo()).Create(id, text)
		fmt.Println("create done🔥")
		todos := handler.NewTodo(infra.NewTodoRepo()).Create(id, text)
		ctx.HTML(200, "index.html", gin.H{"todos": todos})

	})

	// 	router.POST("/update/:id", func(ctx *gin.Context) {
	// 		fmt.Println("PUT now")
	// 		n := ctx.Param("id")
	// 		id, err := strconv.Atoi(n)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		text := ctx.PostForm("text")
	// 		UpdateTodo(id, text)
	// 		ctx.Redirect(302, "/")
	// 	})

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

// func CreateTodo(id int, text string) {
// 	newTodo := entity.Todo{
// 		ID:        id,
// 		Text:      text,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	todos = append(todos, newTodo)
// }

// func UpdateTodo(id int, text string) {
// 	//todosのIDが一致するものを更新する
// 	for i, todo := range todos {
// 		if todo.ID == id {
// 			todos[i].Text = text
// 			todos[i].UpdatedAt = time.Now()
// 		}
// 	}
// }

// // todosのIDが一致するものを削除する
//
//	func DeleteTodo(id int) {
//		for i, todo := range todos {
//			if todo.ID == id {
//				todos = append(todos[:i], todos[i+1:]...)
//			}
//		}
