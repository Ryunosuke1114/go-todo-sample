package handler

import (
	"go-qr-app/domain/model/todo"
	repo "go-qr-app/domain/repository"
)

type Todo struct{ repo repo.Todo }

func NewTodo(repo repo.Todo) *Todo {
	return &Todo{repo: repo}
}

func (t *Todo) Get() []*todo.Todo {
	todos := t.repo.Get()
	return todos
}
