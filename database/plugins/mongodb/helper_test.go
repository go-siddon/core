package mongodb

import (
	"reflect"
	"testing"

	"github.com/go-siddon/core/database"
	"github.com/go-siddon/core/internal/core"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Test_convertParamsToBson(t *testing.T) {
	type args struct {
		filter []database.Params
	}
	tests := []struct {
		name string
		args args
		want bson.D
	}{
		{
			name: "Test Empty Fields",
			args: args{
				filter: []database.Params{},
			},
			want: bson.D{},
		},
		{
			name: "Test One Input",
			args: args{
				filter: []database.Params{
					database.SetParam("name", "jon"),
				},
			},
			want: bson.D{
				{Key: "name", Value: "jon"},
			},
		},
		{
			name: "Test Multiple Input",
			args: args{
				filter: []database.Params{
					database.SetParam("email", "jon@doe.com"),
					database.SetParam("name", "jon"),
				},
			},
			want: bson.D{
				{Key: "email", Value: "jon@doe.com"},
				{Key: "name", Value: "jon"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertParamsToBson(tt.args.filter...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertParamsToBson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertSortParamsToBson(t *testing.T) {
	type args struct {
		filter []database.SortParams
	}
	tests := []struct {
		name string
		args args
		want bson.D
	}{
		{
			name: "Test Empty Fields",
			args: args{
				filter: []database.SortParams{},
			},
			want: bson.D{},
		},
		{
			name: "Test One Input",
			args: args{
				filter: []database.SortParams{
					database.SetSort("name", database.ASC),
				},
			},
			want: bson.D{
				{Key: "name", Value: -1},
			},
		},
		{
			name: "Test Multiple Input",
			args: args{
				filter: []database.SortParams{
					database.SetSort("email", database.ASC),
					database.SetSort("name", database.DESC),
				},
			},
			want: bson.D{
				{Key: "email", Value: -1},
				{Key: "name", Value: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertSortParamsToBson(tt.args.filter...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertSortParamsToBson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertParsedToBSON(t *testing.T) {
	type args struct {
		parsed []core.ParsedField
	}
	tests := []struct {
		name    string
		args    args
		want    bson.D
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertParsedToBSON(tt.args.parsed...)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertParsedToBSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertParsedToBSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIndexes(t *testing.T) {
	type args struct {
		parsed []core.ParsedField
	}
	tests := []struct {
		name    string
		args    args
		want    []mongo.IndexModel
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIndexes(tt.args.parsed...)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIndexes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIndexes() = %v, want %v", got, tt.want)
			}
		})
	}
}
