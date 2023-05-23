package handler

import (
	"fmt"
	"go-qr-app/domain/model/todo"
	"go-qr-app/domain/model/user"
	repo "go-qr-app/domain/repository"
)

type Todo struct{ repo repo.Todo }

func NewTodo(repo repo.Todo) *Todo {
	return &Todo{repo: repo}
}

func (t *Todo) Get() []*todo.Todo {
	todos := t.repo.Get()
	user := user.User{Status: 1}
	return AllowedTodo(int(user.Status), todos)
}

func (t *Todo) GetOne(targetID int) *todo.Todo {
	todo := t.repo.GetOne(targetID)
	return todo
}

func (t *Todo) Create(id int, text string) []*todo.Todo {
	newTodo := t.repo.Create(id, text)
	fmt.Println(newTodo)
	todos := t.repo.Get()
	fmt.Println(todos)
	fmt.Println("アペンド前")
	todos = append(todos, newTodo)
	fmt.Println("アペンド後")

	fmt.Println(todos)
	return todos
}

func AllowedTodo(accessState int, todos []*todo.Todo) []*todo.Todo {
	//渡されたuserの権限に応じて、返すtodoを変える
	var result []*todo.Todo
	switch accessState {
	case 0:
		result = todos
	case 1:
		result = todos
	case 2:
		result = todos
	case 3:
		result = []*todo.Todo{}
	}
	return result
}
