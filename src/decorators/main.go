// https://www.youtube.com/watch?v=9Ia16QOY8rk&index=2&list=PLrCZzMib1e9q-X5V9pTM6J0AemRWseM7I
package main

import (
	"fmt"
)

type iStuff interface {
	DoStuff()
}

type stuff struct {
	iStuff
	Name string
}

// declare method
func (s stuff) SomeComplex() {
	s.DoStuff()
}

type realStuff string

// declare method
func (r realStuff) DoStuff() {
	fmt.Println(r)
}

type fakeStuff int

func (r fakeStuff) DoStuff() {
	fmt.Println("Its a trap!")

}

func main() {
	r := realStuff("Hey!")
	f := fakeStuff(0)

	rS := stuff{r, "stuff"}
	rS.SomeComplex()

	fS := stuff{f, "fake"}
	fS.SomeComplex()

	main2()

}

func main2() {

}
