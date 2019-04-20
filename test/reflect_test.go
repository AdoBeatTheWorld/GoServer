package test

import (
	"reflect"
	"testing"
)

type A struct {
}

func TestReflect(t *testing.T) {
	a := &A{}
	t.Log(reflect.TypeOf(a))
}
