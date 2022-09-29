package main

import (
	"fmt"
)

func F1[T Ordered](x T) {
	fmt.Println(x)
}

func F2(x interface{}) {
	fmt.Println(x)
}

type Ordered interface {
	// integer | Float | ~string
	// Выражение ~string означает набор всех типов, базовым типом которых является string.
	// Это включает в себя сам тип string, а также все типы, объявленные с такими определениями,
	// как type MyString string.
	int | float32 | float64 | ~string
}

func main() {
	F1("ss")
	F1(22)
	// F2("ss")
	// F2(22)
	// fmt.Println(reflect.TypeOf(reflect.Interface))
}
