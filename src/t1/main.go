// http://golang-book.ru/chapter-06-arrays-slices-maps.html
package main

import "fmt"

var (
	// x2 map[string]int
	x2 = make(map[string]int)
)

func main() {
	// var total float64 = 0
	x := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}

	for _, value := range x {
		// total += value
		// fmt.Print(_)
		fmt.Println(value)

	}

	x2["key"] = 10.00

	// fmt.Println(x2["key"])
	// name, ok := x2["Un"]
	// fmt.Println(name, ok)

	if name, ok := x2["key"]; ok {
		fmt.Println(name, ok)
	}

	// if !ok {
	// 	// somehow process this case
	// 	fmt.Println("sss")
	// }
}
