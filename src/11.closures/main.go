package main

import "fmt"

func closureTest() func(x string) (string, bool) {
	fn := func(x string) (string, bool) {
		return x, true
	}

	return fn
}

func main() {
	x := closureTest()
	if res, ok := x("ssss"); ok {
		fmt.Println(res)
	}
}
