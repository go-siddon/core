package mongodb

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-siddon/core/database"
	"github.com/go-siddon/core/internal/core"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func convertParsedToBSON(parsed ...core.ParsedField) (bson.D, error) {
	var res = bson.D{}
	for _, p := range parsed {
		for _, attr := range p.FieldAttrs {
			fmt.Println(attr)
		}
		var val interface{}
		switch p.FieldType.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val = p.FieldValue.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val = p.FieldValue.Uint()
		case reflect.Float32, reflect.Float64:
			val = p.FieldValue.Float()
		case reflect.Bool:
			val = p.FieldValue.Bool()
		case reflect.String:
			val = p.FieldValue.String()
		default:
			return nil, errors.New("unsupported data type provided")
		}
		res = append(res, bson.E{Key: p.FieldTag, Value: val})
	}
	return res, nil
}

func getIndexes(parsed ...core.ParsedField) ([]mongo.IndexModel, error) {
	var model = []mongo.IndexModel{}
	for _, p := range parsed {
		for _, attr := range p.FieldAttrs {
			fmt.Println(attr)
		}
	}
	return model, nil
}

func convertParamsToBson(filter ...database.Params) bson.D {
	output := bson.D{}
	for _, f := range filter {
		output = append(output, bson.E{Key: f.GetKey(), Value: f.GetValue()})
	}
	return output
}

func convertSortParamsToBson(filter ...database.SortParams) bson.D {
	output := bson.D{}
	for _, f := range filter {
		var key int
		switch f.GetValue().String() {
		case database.ASC.String():
			key = -1
		case database.DESC.String():
			key = 1
		default:
			panic(errors.New("invalid sort key provided"))
		}
		output = append(output, bson.E{Key: f.GetKey(), Value: key})
	}
	return output
}
