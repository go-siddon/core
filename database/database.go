package database

import (
	"context"
)

// model(db, table, model).find().one().many().limit()

type Model[T any] interface {
	Find(filter ...Params) Find[T]
	Save() ModelSave[T]
	Update(filter ...Params) ModelUpdate[T]
	Delete(filter ...Params) ModelDelete[T]
}

type Find[T any] interface {
	One() FindOne[T]
	Many() FindMany[T]
}

type FindOne[T any] interface {
	Sort(sort ...SortParams) FindOne[T]
	Column() FindOne[T]
	Exec(ctx context.Context) (*T, error)
}
type FindMany[T any] interface {
	Limit(limit int64) FindMany[T]
	Skip(skip int64) FindMany[T]
	Sort(sort ...SortParams) FindMany[T]
	Column() FindMany[T]
	Exec(ctx context.Context) ([]*T, error)
}

type ModelSave[T any] interface {
	One(data T) Exec[T]
	Many(data ...T) Exec[T]
}

type ModelUpdate[T any] interface {
	One(data T) Exec[T]     //will contain the id to match
	Many(data ...T) Exec[T] //will contain the id to match
}

type ModelDelete[T any] interface {
	One() Exec[T]
	Many() Exec[T]
}

type Exec[T any] interface {
	Exec(ctx context.Context) error
}
