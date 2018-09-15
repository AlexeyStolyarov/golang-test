package main

import (
	"fmt"
	"proj1/conf"
	"sync"
)

func main() {

	wgCounter := 4
	var wg sync.WaitGroup // Создание waitgroup. Исходное значение счётчика - 0
	wg.Add(wgCounter)
	for i := 0; i < wgCounter; i++ {
		go conf.Myfunc(i, &wg)

	}
	//go conf.Myfunc(conf.SiteUrl, &wg)
	//go conf.Myfunc(conf.AbsolutePath, &wg)

	wg.Wait()

	fmt.Println(conf.A)
	fmt.Println(conf.B)
	fmt.Println(conf.D)

	conf.Pointer()
}
