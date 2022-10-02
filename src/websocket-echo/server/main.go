package server

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

func Ws_server(ctx context.Context, wg *sync.WaitGroup) {

	defer wg.Done()

	log.SetFlags(log.LstdFlags | log.LUTC)

	locations := []string{"Home", "Work", "Camphous", "Coworking"}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		traceID := genId()
		log.Printf("%v: Receive request\n", traceID)

		log.Printf("%v: Print header\n", traceID)
		for k := range r.Header {
			log.Printf("%v: header[%v]:%v\n", traceID, k, r.Header.Get(k))
		}
		log.Printf("%v: Read body by 10 bytes\n", traceID)

		chunk := make([]byte, 10)
		for i := 0; true; i++ {
			n, err := r.Body.Read(chunk)
			log.Printf("%v: %v index. read %d bytes '%s' err: %v\n", traceID, i, n, chunk, err)
			if err == io.EOF {
				break
			}
		}
		log.Printf("%v: Finish read body\n", traceID)

		header := w.Header()
		UserNameValue := "tochka"
		if r.Header.Get("User-Name-Data") != "" {
			UserNameValue = "authorizated user name is " + r.Header.Get("User-Name-Data")
		}
		header.Add("User-Id", "12341")
		header.Add("User-Name", UserNameValue)
		header.Add("User-Locations", strings.Join(locations, ";"))

		for _, l := range locations {
			header.Add("Location-"+l, strings.Join([]string{"read", "write"}, ";"))
		}

		w.WriteHeader(http.StatusOK)
	})

	log.Println("Start listening...")
	go http.ListenAndServe(":8080", h)
	<-ctx.Done()

	fmt.Println("Gracefull shtd")
}

func genId() string {
	var b [16]byte
	_, err := rand.Read(b[:])
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b[:])
}
