package database

import (
	"context"
)

// model(db, table, model).find().one().many().limit()

// Database interfaces

type DatabaseAdapter struct {
	Name string
}

type Model[T any] interface {
	Find() Find[T]
	Save() ModelSave[T]
	Update() ModelUpdate[T]
	Delete() ModelDelete[T]
}

type Find[T any] interface {
	One() FindOne[T]
	Many() FindMany[T]
}

type FindOne[T any] interface {
	Sort() FindOne[T]
	Column() FindOne[T]
	Exec(ctx context.Context) (*T, error)
}
type FindMany[T any] interface {
	Limit() FindMany[T]
	Sort() FindMany[T]
	Column() FindMany[T]
	Exec(ctx context.Context) ([]*T, error)
}

type ModelSave[T any] interface {
	One() Save[T]
	Many() Save[T]
}

type Save[T any] interface {
	Exec(ctx context.Context) error
}

type ModelUpdate[T any] interface {
	One() Update[T]  //will contain the id to match
	Many() Update[T] //will contain the id to match
}

type Update[T any] interface {
	Body() Update[T]                // will contain the update body
	Exec(ctx context.Context) error //will contain the update execution
}

type ModelDelete[T any] interface {
	One() Delete[T]
	Many() Delete[T]
}

type Delete[T any] interface {
	Exec(ctx context.Context) error
}
