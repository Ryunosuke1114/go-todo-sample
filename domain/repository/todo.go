package repo

import "go-qr-app/domain/entity"

type TodoRepo interface {
	GET() []*entity.Todo
}
