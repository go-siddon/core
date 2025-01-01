package mongodb

import (
	"context"

	"github.com/go-siddon/core/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Model[T any] struct {
	collection *mongo.Collection
	model      T
}

// RegisterModel takes a connection to the mongo database, the name of the collection as well as the Model.
// It returns the interface that enables communication to the mongodb collection on behalf of the registered model
func RegisterModel[T any](db *Mongo, collection string, model T) database.Model[T] {
	return &Model[T]{
		collection: db.client.Collection(collection),
		model:      model,
	}
}

type modelFind[T any] struct {
	client *mongo.Collection
}

func (mo *Model[T]) Find() database.Find[T] {
	return &modelFind[T]{
		client: mo.collection,
	}
}

type findOne[T any] struct {
	client *mongo.Collection
	filter bson.D
	sort   bson.D
	column bson.D
}

func (mf *modelFind[T]) One() database.FindOne[T] {
	return &findOne[T]{
		client: mf.client,
	}
}

func (fo *findOne[T]) Sort() database.FindOne[T] {
	return fo
}
func (fo *findOne[T]) Column() database.FindOne[T] {
	return fo
}
func (fo *findOne[T]) Exec(ctx context.Context) (*T, error) {
	var res T
	err := fo.client.FindOne(ctx, fo.filter, options.FindOne().
		SetProjection(fo.column).SetSort(fo.sort)).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

type findMany[T any] struct {
	client *mongo.Collection
	filter bson.D
	sort   bson.D
	column bson.D
	limit  int64
	skip   int64
}

func (mf *modelFind[T]) Many() database.FindMany[T] {
	return &findMany[T]{
		client: mf.client,
	}
}

func (fm *findMany[T]) Sort() database.FindMany[T] {
	return fm
}
func (fm *findMany[T]) Limit() database.FindMany[T] {
	return fm
}
func (fm *findMany[T]) Column() database.FindMany[T] {
	return fm
}
func (fm *findMany[T]) Exec(ctx context.Context) ([]*T, error) {
	var res []*T
	result, err := fm.client.Find(ctx, fm.filter, options.Find().
		SetSort(fm.sort).SetSkip(fm.skip).SetLimit(fm.limit).
		SetProjection(fm.column))
	if err != nil {
		return nil, err
	}
	defer result.Close(ctx)
	for result.Next(ctx) {
		var single T
		if err := result.Decode(&single); err != nil {
			return nil, err
		}
		res = append(res, &single)
	}
	return res, nil
}

type modelSave[T any] struct{}

func (mo *Model[T]) Save() database.ModelSave[T] {
	return &modelSave[T]{}
}

type saveOne[T any] struct{}

func (ms *modelSave[T]) One() database.Save[T] {
	return &saveOne[T]{}
}
func (so *saveOne[T]) Exec(ctx context.Context) error {
	return nil
}

type saveMany[T any] struct{}

func (mf *modelSave[T]) Many() database.Save[T] {
	return &saveMany[T]{}
}
func (sm *saveMany[T]) Exec(ctx context.Context) error {
	return nil
}

type modelUpdate[T any] struct{}

func (mo *Model[T]) Update() database.ModelUpdate[T] {
	return &modelUpdate[T]{}
}

type updateOne[T any] struct{}

func (mu *modelUpdate[T]) One() database.Update[T] {
	return &updateOne[T]{}
}
func (uo *updateOne[T]) Body() database.Update[T] {
	return uo
}
func (uo *updateOne[T]) Exec(ctx context.Context) error {
	return nil
}

type updateMany[T any] struct{}

func (mu *modelUpdate[T]) Many() database.Update[T] {
	return &updateMany[T]{}
}
func (um *updateMany[T]) Body() database.Update[T] {
	return um
}
func (um *updateMany[T]) Exec(ctx context.Context) error {
	return nil
}

type modelDelete[T any] struct{}

func (mo *Model[T]) Delete() database.ModelDelete[T] {
	return &modelDelete[T]{}
}

type deleteOne[T any] struct{}

func (md *modelDelete[T]) One() database.Delete[T] {
	return &deleteOne[T]{}
}
func (do *deleteOne[T]) Exec(ctx context.Context) error {
	return nil
}

type deleteMany[T any] struct{}

func (md *modelDelete[T]) Many() database.Delete[T] {
	return &deleteMany[T]{}
}

func (dm *deleteMany[T]) Exec(ctx context.Context) error {
	return nil
}
