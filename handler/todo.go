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
	user := user.User{Status: 0}
	//isValidがtrueの場合のみ、新しいTodoを作成する
	isValid := checkAccessStatus(int(user.Status))
	if isValid {
		newTodo := t.repo.Create(id, text)
		todos := t.repo.Get()
		todos = append(todos, newTodo)
		return todos
	}
	todos := t.repo.Get()
	return todos
}

func (t *Todo) Update(id int, text string) *todo.Todo {
	user := user.User{Status: 0}
	//isValidがtrueの場合のみ、Todoを更新する
	isValid := checkAccessStatus(int(user.Status))
	if isValid {
		todo := t.repo.GetOne(id)
		t.repo.Update(id, text)
		return todo
	}
	fmt.Println("You don't have permission to update")
	return nil
}

func (t *Todo) Delete(id int) {
	user := user.User{Status: 0}
	//isValidがtrueの場合のみ、Todoを削除する
	isValid := checkAccessStatus(int(user.Status))
	if isValid {
		t.repo.Delete(id)
	}
	fmt.Println("You don't have permission to delete")
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

func checkAccessStatus(accessStatus int) bool {
	var isValid bool
	switch accessStatus {
	case 0:
		isValid = true
	case 1:
		isValid = true
	case 2:
		isValid = true
	case 3:
		isValid = false
	}
	return isValid
}
