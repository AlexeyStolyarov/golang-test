package main

import "fmt"

type BaseStr struct {
	descr string
}

type Str struct {
	BaseStr
	id  int
	msg string
}

func (s BaseStr) show() {
	fmt.Println("=========== show =========== ")
	fmt.Println(s.descr)
}

func changeVal(change_str Str, msg string) {
	change_str.msg = msg
}

func changeVal2(change_str *Str, msg string) {
	change_str.msg = msg
}

func main() {
	coral := []string{"blue coral", "staghorn coral", "pillar coral"}
	coral = append(coral, "seahorse")

	fmt.Println("Print slice")
	fmt.Println(coral)

	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}
	fmt.Println("Print map")
	fmt.Println(sammy)

	// Structures
	var st Str
	st.id = 10
	st.msg = "hello"
	st.BaseStr.descr = "desct"

	st.BaseStr.show()
	// var st2 = Str{20, "sss", "sss"}
	fmt.Println(st)

	st3 := Str{
		id:  30,
		msg: "aaa",
	}

	// The st's variable type: main.Str
	fmt.Printf("\nThe st's variable type: %T\n", st3)

	fmt.Println(st.id)
	fmt.Println(st3)

	// doesn't change
	changeVal(st3, "HHHHHHHHH")
	fmt.Println(st3)

	// changes
	changeVal2(&st3, "HHHHHHHHH")
	fmt.Println(st3)

	// это будет указатель на созданный объект:
	var str_pointer = new(Str)

	// The st's variable type: *main.Str
	fmt.Printf("\nThe st's variable type: %T\n", str_pointer)

	str_pointer.id = 42
	changeVal2(str_pointer, "42")
	fmt.Println(str_pointer.msg)

}
