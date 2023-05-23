package infra

import (
	"go-qr-app/domain/model/todo"
	repo "go-qr-app/domain/repository"

	"time"
)

type Todo struct {
	todos []*todo.Todo
}

func NewTodoRepo() repo.Todo {
	return &Todo{}
}

func (*Todo) Get() []*todo.Todo {
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
	return todos
}

func (*Todo) GetOne(targetID int) *todo.Todo {
	var todos = []*todo.Todo{
		{
			ID:         1,
			Text:       "test1",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			DeleatedAt: time.Time{},
		},
		{
			ID: 2, Text: "test2",
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
