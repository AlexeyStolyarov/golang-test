package conf

import "fmt"

func Pointer() {
	i := 5
	z := &i

	fmt.Println("i:=", i)

	*z = 10
	fmt.Println("i:=", i)

}
