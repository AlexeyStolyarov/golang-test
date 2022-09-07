// https://habr.com/ru/post/490336/
package main

import (
	"fmt"
	"sync"
	"time"
)

func service(wg *sync.WaitGroup, instance int, wait int) {
	time.Sleep(time.Duration(wait) * time.Second)
	fmt.Println("Service called on instance", instance)
	wg.Done() // decrement counter
}

func main() {
	fmt.Println("main() started")
	var wg sync.WaitGroup // create waitgroup (empty struct)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go service(&wg, i, 1)
	}

	wg.Add(1)
	go service(&wg, 20, 20)

	fmt.Println("wait till the last go routine finished")
	wg.Wait() // blocks here
	fmt.Println("main() stopped")
}
