package mongodb

import (
	"errors"

	"github.com/go-siddon/core/database"
	"github.com/go-siddon/core/internal/core"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func convertParsedToBSON([]core.ParsedField) bson.D {
	return nil
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
			key = 1
		case database.DESC.String():
			key = -1
		default:
			panic(errors.New("invalid sort key provided"))
		}
		output = append(output, bson.E{Key: f.GetKey(), Value: key})
	}
	return output
}
