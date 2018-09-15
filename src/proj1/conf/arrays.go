package conf

import (
	"fmt"
)

var (
	thisIsArray  [3]int
	thisIsArray2 = [...]int{1, 2, 3}
)

// we cant use these arrays in this function because length mismatch

func f1(x [4]int) {
	fmt.Print(x)
}

// and thus we need slices.

var (
	thisIsSlice = make([]int, 3)  // filled slice
)


// but beware go outside border.
// print(thisIsSlice[0]) // error

thisIsSlice = append(thisIsSlice,10)
print(thisIsSlice[0]) // now OK

// добавление слайса в слайс/
// s1 = append(s1, s2)  не работает
// используй ... (sic!)
// s1 = append(s1, s2...)

func f2(){
	slm:= [][]int
	slm = append(slm,s1) //works. because slm is slice type of []int
}



// map aka hash aka associate array.
var(
	myMap map[string]string
	myMap2 map[string]int
	//         index  type
	mapOfMaps = map[string] map[string]string 
	// by default map has nil value. and we cannot use it.
	// thus we have to fill map.
	myMap3 map[string]int = map[string]int{}
	// or myMap3 := map[string]int{}
	myMap4 = make(map[string]int)
	
	

)
