package infra

import (
	"go-qr-app/domain/entity"
	"math/rand"

	"time"
)

type TodoRepo struct {
	todos []*entity.Todo
}

func NewTodoRepo() *TodoRepo {
	return &TodoRepo{}
}

func (t *TodoRepo) GET() []*entity.Todo {
	var todos = []*entity.Todo{
		{
			ID:         rand.Int(),
			Text:       "test1",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			DeleatedAt: time.Time{},
		},
		{ID: rand.Int(), Text: "test2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: rand.Int(), Text: "test3", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: rand.Int(), Text: "test4", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	return todos
}
