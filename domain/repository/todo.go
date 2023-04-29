package repo

import "go-qr-app/domain/model/todo"

type Todo interface {
	Get() []*todo.Todo
}
