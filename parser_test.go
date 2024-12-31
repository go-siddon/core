package core

import (
	"reflect"
	"testing"
)

func Test_parser_Parse(t *testing.T) {
	testOne := struct {
		Name string `db:"name" attr:"required"`
	}{
		Name: "Jon Doe",
	}
	type args struct {
		val interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []ParsedField
		wantErr bool
	}{
		{
			name: "Check parser",
			args: args{
				val: testOne,
			},
			want: []ParsedField{
				{
					FieldAttrs: []string{"required"},
					FieldValue: reflect.ValueOf(testOne).Field(0),
					FieldType:  reflect.ValueOf(testOne).Type().Field(0).Type,
					FieldName:  reflect.ValueOf(testOne).Type().Field(0).Name,
					FieldDB:    "name",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{}
			got, err := p.Parse(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !compareParserResponse(got, tt.want) {
				t.Errorf("parser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
