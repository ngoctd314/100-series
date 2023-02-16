package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// pprof: go tool pprof -http localhost:9000 http://localhost:8080/debug/pprof/heap
	go consumeMessages()

	http.ListenAndServe(":8080", nil)
}

func consumeMessages() {
	for {
		msg := receiveMessage()
		// Do something with msg
		storeHeader(getHeader(msg))
		log.Println("len header", len(_headers))
	}
}
func receiveMessage() []byte {
	return make([]byte, 2<<20)
}

func getHeader(msg []byte) []byte {
	// header := make([]byte, 5)
	// n := copy(header, msg)
	// log.Printf("copy %d bytes\n", n)

	header := msg[:5:5]

	return header
}

var (
	_headers [][]byte
)

func storeHeader(header []byte) {
	_headers = append(_headers, header)
}
