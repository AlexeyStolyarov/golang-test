package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ch1 := make(chan string)
	timeout := time.After(3 * time.Second)
	pollInt := time.Second
	// var quit = make(chan struct{})
	ch_num := 10

	for i := 0; i < ch_num; i++ {
		go gr1(ch1, 10-i, i)
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
		case <-timeout:
			fmt.Println("my patience ran out")
			return // Сдается и возвращается
		default:
			fmt.Println("waiting")
			// time.Sleep(500 * time.Millisecond)
		}
		time.Sleep(pollInt)
	}
}

func gr1(ch chan string, sleep int, arg int) {
	time.Sleep(time.Duration(sleep) * time.Second)
	fmt.Println("test: " + strconv.Itoa(arg))
	ch <- strconv.Itoa(arg)

	// close(ch)
}
