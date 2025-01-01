package mongodb

import (
	"context"

	"github.com/go-siddon/core/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Model[T any] struct {
	collection *mongo.Collection
	model      T
	database.Model[T]
}

// RegisterModel takes a connection to the mongo database, the name of the collection as well as the Model.
// It returns the interface that enables communication to the mongodb collection on behalf of the registered model
func RegisterModel[T any](db *mongo.Database, collection string, model T) *Model[T] {
	return &Model[T]{
		collection: db.Collection(collection),
		model:      model,
	}
}

type modelFind struct{}

func (mo *Model[T]) Find() *modelFind {
	return &modelFind{}
}

func (mf *modelFind) One()  {}
func (mf *modelFind) Many() {}

type modelSave struct{}

func (mo *Model[T]) Save() *modelSave {
	return &modelSave{}
}
func (mf *modelSave) One()  {}
func (mf *modelSave) Many() {}

type modelUpdate struct{}

func (mo *Model[T]) Update() *modelUpdate {
	return &modelUpdate{}
}

type updateOne struct{}
type updateMany struct{}

func (mf *modelUpdate) One() *updateOne {
	return &updateOne{}
}
func (mf *modelUpdate) Many() *updateMany {
	return &updateMany{}
}
func (uo *updateOne) Exec(ctx context.Context) error {
	return nil
}
func (um *updateMany) Exec(ctx context.Context) error {
	return nil
}

type modelDelete struct{}

func (mo *Model[T]) Delete() *modelDelete {
	return &modelDelete{}
}

type deleteOne struct{}
type deleteMany struct{}

func (mf *modelDelete) One() *deleteOne {
	return &deleteOne{}
}
func (mf *modelDelete) Many() *deleteMany {
	return &deleteMany{}
}

func (do *deleteOne) Exec(ctx context.Context) error {
	return nil
}
func (dm *deleteMany) Exec(ctx context.Context) error {
	return nil
}
