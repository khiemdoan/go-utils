package memoryutils

import (
	"reflect"
	"testing"
)

func TestNewObject(t *testing.T) {
	type abc struct {
		number int
		text   string
	}

	x := abc{
		number: 12,
		text:   "abcxyz",
	}

	y := NewObject(x)

	if reflect.TypeOf(y) != reflect.TypeOf(x) {
		t.Error("NewObject was incorrect")
	}

	if _, ok := y.(abc); !ok {
		t.Error("NewObject was incorrect")
	}
}
