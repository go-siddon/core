package mongodb

import (
	"context"

	"github.com/go-siddon/core/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Model[T any] struct {
	client *mongo.Collection
	model  T
}

// RegisterModel takes a connection to the mongo database, the name of the collection as well as the Model.
// It returns the interface that enables communication to the mongodb collection on behalf of the registered model
func RegisterModel[T any](db *Mongo, collection string, model T) database.Model[T] {
	return &Model[T]{
		client: db.client.Collection(collection),
		model:  model,
	}
}

type modelFind[T any] struct {
	client *mongo.Collection
	filter bson.D
}

func (mo *Model[T]) Find(filter ...database.Params) database.Find[T] {
	fil := convertParamsToBson(filter...)
	return &modelFind[T]{
		client: mo.client,
		filter: fil,
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
		filter: mf.filter,
	}
}

func (fo *findOne[T]) Sort(sortkeys ...database.SortParams) database.FindOne[T] {
	fo.sort = convertSortParamsToBson(sortkeys...)
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
		filter: mf.filter,
	}
}

func (fm *findMany[T]) Sort(sortkey ...database.SortParams) database.FindMany[T] {
	fm.sort = convertSortParamsToBson(sortkey...)
	return fm
}
func (fm *findMany[T]) Limit(limit int64) database.FindMany[T] {
	fm.limit = limit
	return fm
}
func (fm *findMany[T]) Skip(skip int64) database.FindMany[T] {
	fm.skip = skip
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

type modelSave[T any] struct {
	client *mongo.Collection
}

func (mo *Model[T]) Save() database.ModelSave[T] {
	return &modelSave[T]{
		client: mo.client,
	}
}

type saveOne[T any] struct {
	client   *mongo.Collection
	document T
}

func (ms *modelSave[T]) One() database.Exec[T] {
	return &saveOne[T]{
		client: ms.client,
	}
}
func (so *saveOne[T]) Exec(ctx context.Context) error {

	if _, err := so.client.InsertOne(ctx, nil); err != nil {
		return err
	}

	return nil
}

type saveMany[T any] struct {
	client   *mongo.Collection
	document bson.D
}

func (ms *modelSave[T]) Many() database.Exec[T] {
	return &saveMany[T]{
		client: ms.client,
	}
}
func (sm *saveMany[T]) Exec(ctx context.Context) error {
	if _, err := sm.client.InsertMany(ctx, nil); err != nil {
		return err
	}
	return nil
}

type modelUpdate[T any] struct {
	client *mongo.Collection
}

func (mo *Model[T]) Update(filter ...database.Params) database.ModelUpdate[T] {
	return &modelUpdate[T]{
		client: mo.client,
	}
}

type updateOne[T any] struct {
	client   *mongo.Collection
	filter   bson.D
	document bson.D
}

func (mu *modelUpdate[T]) One() database.Exec[T] {
	return &updateOne[T]{
		client: mu.client,
	}
}

func (uo *updateOne[T]) Exec(ctx context.Context) error {
	if _, err := uo.client.UpdateOne(ctx, uo.filter, uo.document); err != nil {
		return err
	}
	return nil
}

type updateMany[T any] struct {
	client   *mongo.Collection
	filter   bson.D
	document bson.D
}

func (mu *modelUpdate[T]) Many() database.Exec[T] {
	return &updateMany[T]{
		client: mu.client,
	}
}

func (um *updateMany[T]) Exec(ctx context.Context) error {
	if _, err := um.client.UpdateMany(ctx, um.filter, um.document); err != nil {
		return err
	}
	return nil
}

type modelDelete[T any] struct {
	client *mongo.Collection
}

func (mo *Model[T]) Delete(filter ...database.Params) database.ModelDelete[T] {
	return &modelDelete[T]{
		client: mo.client,
	}
}

type deleteOne[T any] struct {
	client *mongo.Collection
	filter bson.D
}

func (md *modelDelete[T]) One() database.Exec[T] {
	return &deleteOne[T]{
		client: md.client,
	}
}
func (do *deleteOne[T]) Exec(ctx context.Context) error {
	if _, err := do.client.DeleteOne(ctx, do.filter); err != nil {
		return err
	}
	return nil
}

type deleteMany[T any] struct {
	client *mongo.Collection
	filter bson.D
}

func (md *modelDelete[T]) Many() database.Exec[T] {
	return &deleteMany[T]{
		client: md.client,
	}
}

func (dm *deleteMany[T]) Exec(ctx context.Context) error {
	if _, err := dm.client.DeleteMany(ctx, dm.filter); err != nil {
		return err
	}
	return nil
}
