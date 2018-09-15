package conf

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	// A iota - autoincrement. A,B,D => 1,2,4
	A = iota + 1
	B
	_
	D
)

var (
	SiteUrl      string
	AbsolutePath string
)

func init() {
	SiteUrl = "http://loctalk.net"
	AbsolutePath = "/path/to/project/"
	/*
		mode := os.Getenv("MARTINI_ENV")

		switch mode {
		case "production":
			SiteUrl = "http://loctalk.net"
			AbsolutePath = "/path/to/project/"
		default:
			SiteUrl = "http://127.0.0.1"
			AbsolutePath = "/path/to/project/"
		} */

}

func Myfunc(x int, extWg *sync.WaitGroup) {
	defer extWg.Done()

	time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
	fmt.Println(x)
}
