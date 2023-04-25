package repo

import "go-qr-app/domain/entity"

type Todo interface {
	Get() []*entity.Todo
}
