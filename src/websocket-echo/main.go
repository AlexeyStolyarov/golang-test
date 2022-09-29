package main

import (
	"fmt"

	"websocket-echo/config"
)

func main() {
	var c config.Conf

	c.GetConf(config.File)

	fmt.Printf("%+v\n", c)
}
