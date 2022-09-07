package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

var (
	myRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	orders = make(chan string)
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// rand.Seed(time.Now().UnixNano())
	taxi := []string{
		"yandex",
		"deli",
		"uber",
	}

	for _, i := range taxi {
		// fmt.Println(i)
		go call(i)
	}

	for {
		select {
		case order, ok := <-orders:
			if ok {
				fmt.Println(&ctx, order)
			}
		default:
			time.Sleep(time.Second)
		}
	}
}

func call(ctx *context.WithCancel t string) {
	time.Sleep(time.Duration(myRand.Intn(10)) * time.Second)
	fmt.Println("calling for " + t)
	orders <- t
}
