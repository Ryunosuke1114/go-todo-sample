package handler

import (
	"go-qr-app/domain/entity"
	repo "go-qr-app/domain/repository"
)

type Todo struct{ repo repo.Todo }

func NewTodo(repo repo.Todo) *Todo {
	return &Todo{repo: repo}
}

func (t *Todo) Get() []*entity.Todo {
	todos := t.repo.Get()
	return todos
}
