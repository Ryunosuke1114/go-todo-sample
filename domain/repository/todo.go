package repo

import "go-qr-app/domain/entity"

type TodoRepo interface {
	Get() []*entity.Todo
}
