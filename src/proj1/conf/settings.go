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

const (
	_         = iota
	KB uint64 = 1 << (10 * iota)
	MB
	GB
)

var (
	SiteUrl      string
	AbsolutePath string
)

func init() {
	SiteUrl = "http://loctalk.net"
	AbsolutePath = "/path/to/project/"

	fmt.Println(KB)
	fmt.Println(MB, GB)
}

func Myfunc(x int, extWg *sync.WaitGroup) {
	defer extWg.Done()

	time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
	fmt.Println(x)
}
