// package Parser contains the implementation of the parser that marshals/converts a struct to a readable format
// that other pacakges can use

// it expects the struct to have a tag of  `db:"<field-name>"` and or `attr:"<attribute-name>"` in order to
// parse them correctly.

// Example:

// type User struct {
// 	ID      string    `db:"id" attr:"required,mongoid"`
// 	Name    string    `db:"name" attr:"min=2,max=20"`
// 	Email   string    `db:"email" attr:"email,required"`
// 	Age     int       `db:"age" attr:"required"`
// 	DOB     time.Time `db:"date_of_birth" attr:"required,tz"`
// 	Balance float32   `db:"balance"`
// 	Street  Address   `attr:"embed"`
// 	Friends []Friends `db:"friends"`
// }

// type Address struct {
// 	Street  string `db:"street"`
// 	State   string `db:"state"`
// 	ZipCode string `db:"zip_code"`
// 	Country string `db:"country"`
// }

// type Friends struct {
// 	Name  string `db:"name"`
// 	Email string `db:"email" attr:"email"`
// 	Phone uint   `db:"phone" attr:"required"`
// }

package core

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const (
	databaseTag = "db"

	attributeTag = "attr"

	skipFieldTag = "-"
)

type attrTags int

const (
	requiredTag attrTags = iota
	minTag
	maxTag
	emailTag
	mongoidTag
	tzTag
)

var (
	ErrNotStruct = errors.New("error: invalid type received. expected type %s, but got type %s")
)

type parser struct{}

func New() *parser {
	return &parser{}
}

// ParserField is contains information about the fields in the struct that was parsed,
// information pertaining to its value, type, name, structtags etc
type ParsedField struct {
	FieldAttrs []string
	FieldValue reflect.Value
	FieldType  reflect.Type
	FieldName  string
	FieldDB    string
}

func (p *parser) Parse(obj interface{}) ([]ParsedField, error) {
	var sliceOfParsed []ParsedField
	v := reflect.ValueOf(obj)

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf(ErrNotStruct.Error(), reflect.Struct.String(), v.Kind().String())
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		var single ParsedField
		single.FieldType = field.Type
		single.FieldName = field.Name
		single.FieldValue = v.Field(i)

		tag := field.Tag

		if def, ok := tag.Lookup(attributeTag); ok {
			attr := strings.Split(def, ",")
			single.FieldAttrs = attr
		}

		if def, ok := tag.Lookup(databaseTag); ok {
			single.FieldDB = def
		}
		sliceOfParsed = append(sliceOfParsed, single)
	}

	return sliceOfParsed, nil
}

func compareParserResponse(a, b []ParsedField) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i].FieldName != b[i].FieldName ||
			a[i].FieldDB != b[i].FieldDB ||
			!reflect.DeepEqual(a[i].FieldAttrs, b[i].FieldAttrs) ||
			!reflect.DeepEqual(a[i].FieldValue.Interface(), b[i].FieldValue.Interface()) {
			return false
		}
	}
	return true
}
