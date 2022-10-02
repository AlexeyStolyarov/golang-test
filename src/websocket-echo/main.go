package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"websocket-echo/server"
)

func main() {
	// var c config.Conf
	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	// c.GetConf(config.File)

	// fmt.Printf("%+v\n", c)

	wg.Add(1)
	go server.Ws_server(ctx, &wg)

	wg.Wait()
}
