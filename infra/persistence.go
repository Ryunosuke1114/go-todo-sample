package infra

import (
	"go-qr-app/domain/model/todo"
	repo "go-qr-app/domain/repository"

	"time"
)

// 全体で共通の変数を定義する
var todos = []*todo.Todo{
	{
		ID:         1,
		Text:       "test1",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeleatedAt: time.Time{},
	},
	{
		ID:        2,
		Text:      "test2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        3,
		Text:      "test3",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        4,
		Text:      "test4",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

type Todo struct {
	todos []*todo.Todo
}

func NewTodoRepo() repo.Todo {
	return &Todo{}
}

func (*Todo) Get() []*todo.Todo {
	return todos
}

func (*Todo) GetOne(targetID int) *todo.Todo {
	for _, todo := range todos {
		if todo.ID == targetID {
			return todo
		}
	}
	return nil
}

func (*Todo) Create(id int, text string) *todo.Todo {
	newTodo := &todo.Todo{
		ID:        id,
		Text:      text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return newTodo
}

func (*Todo) Update(id int, text string) {
	for _, todo := range todos {
		if todo.ID == id {
			todo.Text = text
			todo.UpdatedAt = time.Now()
		}
	}
}

func (*Todo) Delete(id int) {
	//新しい配列を作成し、削除対象のTodo以外を新しい配列に追加する
	var newTodos []*todo.Todo
	for _, todo := range todos {
		if todo.ID != id {
			newTodos = append(newTodos, todo)
		}
	}
	todos = newTodos
}
