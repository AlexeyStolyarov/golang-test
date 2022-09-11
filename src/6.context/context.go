package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	myRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func main() {
	var (
		ordersCh = make(chan string)
		taxi     = []string{"yandex", "deli", "uber", "SberTaxi"}
		winner   string
		wg       sync.WaitGroup
	)

	ctx, cancel := context.WithCancel(context.Background())
	// timeout := time.After(10 * time.Second)

	defer cancel()
	// rand.Seed(time.Now().UnixNano())

	// doesn't work in that way!!
	// for _, i := range taxi {
	// fmt.Println(i)
	for i := range taxi {
		tt := taxi[i]

		wg.Add(1)

		go func() {
			// doesn't work in that way!!
			// go call(ctx, tt, ordersCh)
			call(ctx, tt, ordersCh)
			wg.Done()
		}()
	}

	// cancel()
	go func() {
		winner = <-ordersCh
		cancel()
	}()

	wg.Wait()
	log.Printf("found car in %q", winner)
	fmt.Println("main: quit!!")

}

func call(ctx context.Context, t string, ordersCh chan string) {
	time.Sleep(2 * time.Second)

	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println(t + "says: timeout!")
			return
		case <-ctx.Done():
			fmt.Println(t + " says: bye!! ->" + ctx.Err().Error())
			return
		default:
			time.Sleep(time.Millisecond)
			if rand.Float64() > 0.75 {
				ordersCh <- t
				return
			}
		}

	}

}
