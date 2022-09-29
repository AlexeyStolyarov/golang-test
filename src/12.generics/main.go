package main

import (
	"fmt"

	// go get -u golang.org/x/exp/constraints
	// go mod build
	// go mod vendor
	"golang.org/x/exp/constraints"
)

// func myFunc[T comparable](x T) {
// func myFunc[T any](x T) {
func myFunc[T constraints.Ordered](x T) {
	fmt.Println(x)
}

func main() {

	x := "string"
	y := 42
	// z := []string{"11", "bb", "cc"}
	myFunc(x)
	myFunc(y)
	// myFunc(z)

}
