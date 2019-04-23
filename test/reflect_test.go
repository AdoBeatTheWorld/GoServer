package test

import (
	"fmt"
	"reflect"
	"testing"
)

type FooIF interface {
	DoSomething()
	DoSomethingWithArg(a string)
	DoSomethingWithUnCertenArg(a ...string)
}

type Foo struct {
	A int
	B string
	c struct {
		C1 int
	}
}

func (f *Foo) DoSomthing() {
	fmt.Println(f.A, f.B)
}

func (f *Foo) DoSomethingWithArg(a string) {
	fmt.Println(f.A, f.B, a)
}

func (f *Foo) DoSomthingWithUncentenArg(a ...string) {
	fmt.Println(f.A, f.B, a[0])
}

func (f *Foo) returnOneResult() int {
	return 2
}

func TestReflect(t *testing.T) {
	var simpleObj Foo
	var pointer2obj = &simpleObj
	var simpleIntArray = [3]int{1, 2, 3}
	var simpleMap = map[string]string{
		"a": "b",
	}
	var simpleChan = make(chan int, 1)
	var x uint64
	var y uint32

	varType := reflect.TypeOf(simpleObj)
	varPointerType := reflect.TypeOf(pointer2obj)

	fmt.Println("Align:", varType.Align())
	fmt.Println("FieldAlign:", varType.FieldAlign())
	fmt.Println("Name:", varType.Name())
	fmt.Println("PkgPath:", varType.PkgPath())
	fmt.Println("Size:", varType.Size())
	fmt.Println("Kind:", varType.Kind())
	fmt.Println("NumMethod:", varPointerType.NumMethod())

	m, success := varPointerType.MethodByName("DoSomethingWithArg")
	if success {
		m.Func.Call([]reflect.Value{
			reflect.ValueOf(pointer2obj),
			reflect.ValueOf("sad"),
		})
	}

	m = varPointerType.Method(0)
	m.Func.Call([]reflect.Value{
		reflect.ValueOf(pointer2obj),
		reflect.ValueOf("sad2"),
	})

	fmt.Println("Implements:", varPointerType.Implements(reflect.TypeOf((*FooIF)(nil)).Elem()))
	fmt.Println("Bits:", reflect.TypeOf(x).Bits())
	fmt.Println("Elem:", reflect.TypeOf(simpleIntArray).Elem().Kind())
	fmt.Println("Len:", reflect.TypeOf(simpleIntArray).Len())
	fmt.Println("Field:", varType.Field(1))
	fmt.Println("FieldByIndex:", varType.FieldByIndex([]int{2, 0}))

	fi, success2 := varType.FieldByName("A")
	if success2 {
		fmt.Println("FieldByName:", fi)
	}

	fi, success2 = varType.FieldByNameFunc(func(fieldName string) bool {
		return fieldName == "A"
	})

	if success2 {
		fmt.Println("FieldByName:", fi)
	}

	fmt.Println("NumField:", varType.NumField())
	fmt.Println("Key:", reflect.TypeOf(simpleMap).Key().Name())
	fmt.Println("NumIn:", reflect.TypeOf(pointer2obj.DoSomthingWithUncentenArg).NumIn())
	fmt.Println("In:", reflect.TypeOf(pointer2obj.DoSomthingWithUncentenArg).In(0))
	fmt.Println("IsVariadic:", reflect.TypeOf(pointer2obj.DoSomthingWithUncentenArg).IsVariadic())
	fmt.Println("NumOut:", reflect.TypeOf(pointer2obj.DoSomthingWithUncentenArg).NumOut())
	fmt.Println("Out:", reflect.TypeOf(pointer2obj.returnOneResult).Out(0))
	fmt.Println("ChanDir:", int(reflect.TypeOf(simpleChan).ChanDir()))
	fmt.Println("Comparabld:", varPointerType.Comparable())
	fmt.Println("ConvertibleTo:", varPointerType.ConvertibleTo(reflect.TypeOf("a")))
	fmt.Println("AssignableTo:", reflect.TypeOf(x).AssignableTo(reflect.TypeOf(y)))
	var x1 float64 = 12.3
	v := reflect.ValueOf(x1)
	fmt.Println("Canset:", v.CanSet())
}
