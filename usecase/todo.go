package usecase

import "go-qr-app/domain/model/todo"

func TodoFilter(accessState int, todos []*todo.Todo) []*todo.Todo {
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
