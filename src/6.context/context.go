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
	taxi   = []string{
		"yandex",
		"deli",
		"uber",
	}
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// timeout := time.After(10 * time.Second)

	defer cancel()
	// rand.Seed(time.Now().UnixNano())

	for _, i := range taxi {
		// fmt.Println(i)
		go call(ctx, i)
	}

	// cancel()

	order := <-orders
	fmt.Println("main: Got order from: " + order)
	cancel()

	fmt.Println("main: quit!!")

}

func call(ctx context.Context, t string) {
	defer fmt.Println(t + " says: bye!!")

	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println(t + "says: timeout!")
			return
		case <-ctx.Done():
			fmt.Print(ctx.Err())
		default:
			time.Sleep(time.Duration(myRand.Intn(10)) * time.Second)
			fmt.Println("calling for " + t)
			orders <- t

		}

	}

}
