// https://github.com/luciotato/golang-notes/blob/master/OOP.md
package main

import (
	"fmt"
	"strconv"
)

type BaseClass struct {
	name string
	age  int
	// sex  string
}

type Man struct {
	BaseClass
	money int
}

// type Woman struct {
// 	BaseClass
// 	tits int
// }

func (s BaseClass) sayMyName() {
	fmt.Println("My name is " + s.name)
}

func (s BaseClass) sayMyAge() {
	fmt.Println("BaseClass>>>" + strconv.Itoa(s.age))
}

func (s Man) sayMyAge() {
	fmt.Println("Man>>>" + strconv.Itoa(s.age))
}

func main() {
	john := Man{
		money: 100,
		BaseClass: BaseClass{
			age:  42,
			name: "John",
		},
	}

	// they are the same
	// sayMyName wasn't shadowed
	john.sayMyName()
	john.BaseClass.sayMyName()

	// BaseClass.sayMyAge() is shadowed
	john.sayMyAge()
	john.BaseClass.sayMyAge()
}
