package reflectutils

import (
	"reflect"
)

func NewObject(original any) any {
	var reflected any
	typeOf := reflect.TypeOf(original)
	if typeOf.Kind() == reflect.Ptr {
		reflected = reflect.New(reflect.ValueOf(original).Elem().Type()).Interface()
	} else {
		reflected = reflect.New(typeOf).Elem().Interface()
	}
	return reflected
}
