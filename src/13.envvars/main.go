package main

import (
	"fmt"
	"os"
)

func main() {

	getEnv := func(key string) {
		val, ok := os.LookupEnv(key)
		if !ok {
			fmt.Printf("%s not set\n", key)
		} else {
			fmt.Printf("%s=%s\n", key, val)
		}
	}

	getEnv("EDITOR")
	getEnv("SHELL")
}
