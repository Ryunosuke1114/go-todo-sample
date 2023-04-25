package todo

import (
	"time"
)

type Todo struct {
	ID         int
	Text       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleatedAt time.Time
}

//何かしらのバリデーション
// func validate(){

// }
