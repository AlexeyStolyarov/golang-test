package conf

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
const (
	SITE_NAME      string = "LocTalk"
	DEFAULT_LIMIT  int    = 10
	MAX_LIMIT      int    = 1000
	MAX_POST_CHARS int    = 1000
)
*/

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
