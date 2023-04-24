package main

import (
	"go-qr-app/infra"
	"strconv"

	// "math/rand"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	repo := infra.NewTodoRepo()

	router.GET("/", func(ctx *gin.Context) {
		todos := repo.Get()
		ctx.HTML(200, "index.html", gin.H{"todos": todos})
	})

	router.GET("/detail/:id", func(ctx *gin.Context) {
		p := ctx.Param("id")
		id, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		var todo = repo.GetOne(id)
		ctx.HTML(200, "detail.html", gin.H{"todo": todo})
	})

	// 	router.POST("/new", func(ctx *gin.Context) {
	// 		id := rand.Int()
	// 		text := ctx.PostForm("text")
	// 		CreateTodo(id, text)
	// 		ctx.Redirect(302, "/")
	// 	})

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
