package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var ch1 = make(chan string)
	ch_num := 10

	for i := 0; i < ch_num; i++ {
		go gr1(ch1, 10-i, "test"+strconv.Itoa(i))
	}

	// for i := 0; i < ch_num; i++ {
	// 	fmt.Println(i)
	// 	fmt.Println(<-ch1)
	// }

	for {
		select {
		case res, ok := <-ch1:
			if ok {
				fmt.Println(res)
			}
		default:
			// fmt.Println("waiting")
			// time.Sleep(500 * time.Millisecond)

		}
	}
}

func gr1(ch chan string, sleep int, arg string) {
	time.Sleep(time.Duration(sleep) * time.Second)
	// fmt.Println("test: " + arg)
	ch <- arg

	// close(ch)
}
