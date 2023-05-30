package repo

import "go-qr-app/domain/model/todo"

type Todo interface {
	Get() []*todo.Todo
	GetOne(targetID int) *todo.Todo
	Create(id int, text string) *todo.Todo
	Update(id int, text string)
}
