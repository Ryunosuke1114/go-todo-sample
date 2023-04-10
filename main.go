package main

import (
	"go-qr-app/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	dbInit()

	router.GET("/", func(ctx *gin.Context) {
		todos := dbGetAll()
		ctx.HTML(200, "index.html", gin.H{"todos": todos})
	})

	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dnInsert(text, status)
		ctx.Redirect(302, "/")
	})

	router.GET("/detail/*id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n) //intに変換
		if err != nil {
			panic(err)
		}
		todo := dbGetOne(id)
		ctx.HTML(200, "detail HTML", gin.H{"todo": todo})

		router.POST("/update/:id", func(ctx *gin.Context) {
			n := ctx.Param("id")
			id, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			text := ctx.PostForm("text")
			status := ctx.PostForm("status")
			dbUpdate(id, text, status)
			ctx.Redirect(302, "/")
		})

	})

	//削除確認
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := dbGetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	})

	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		dbDeleta(id)
		ctx.Redirect(302, "/")
	})
	router.Run()
}

// DB初期化
func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("開かない")
	}
	db.AutoMigrate(&model.Todo{})
	defer db.Close()
}

// DBインサート
func dnInsert(text string, status string) {
	db, err := gorm.Open("sqlite", "text.sqlite3")
	if err != nil {
		panic("開かない")
	}
	db.Create(&model.Todo{Text: text, Status: status})
	defer db.Close()
}

// DBアップデート
func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("開かない")
	}
	var todo model.Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

// DBデリート
func dbDeleta(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("開かない")
	}
	var todo model.Todo
	db.First(&todo.ID)
	db.Delete(&todo)
	db.Close()
}

// DB全件取得
func dbGetAll() []model.Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("開かない")
	}
	var todos []model.Todo
	db.Order("create_at desc").Find(&todos)
	db.Close()
	return todos
}

// DB一件取得
func dbGetOne(id int) model.Todo {
	db, err := gorm.Open("sqlite", "test.sqlite3")
	if err != nil {
		panic("開かない")
	}
	var todo model.Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
