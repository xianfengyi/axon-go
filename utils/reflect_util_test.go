package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type personStruct struct {
	name string
	age  int
}

type testStruct struct {
	id     int64
	person personStruct
}

func TestGetFieldValue(t *testing.T) {
	testStructCase := testStruct{
		id: 123,
		person: personStruct{
			name: "pioneeryi",
			age:  26,
		},
	}
	printTypes(reflect.TypeOf(&testStructCase).Elem())

	vv := reflect.Indirect(reflect.ValueOf(testStructCase))

	person := vv.FieldByName("person")

	assert.Equal(t, int64(123), vv.FieldByName("id").Int())
	assert.Equal(t, "pioneeryi", person.FieldByName("name").String())
}

func printTypes(t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind() == reflect.Struct {
			printTypes(field.Type)
			continue
		}

		column := field.Name
		fmt.Println("column: ", column, field.Type.Name())
	}
}
