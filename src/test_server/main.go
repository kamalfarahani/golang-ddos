package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var reqNum int

func main() {
	http.HandleFunc("/reqtest", getReqTest)

	http.ListenAndServe(":8080", nil)
}

func getReqTest(rw http.ResponseWriter, req *http.Request) {
	reqNum++
	fmt.Println(reqNum)

	rw.Header().Add("Content-Type", "text/plain")
	rw.Write([]byte(strconv.Itoa(reqNum)))
}
