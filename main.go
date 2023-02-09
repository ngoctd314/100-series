package main

import (
	"log"
)

var a = func() error {
	log.Println("initialize a")
	return nil
}()

func main() {
}
