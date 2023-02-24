package main

import (
	"fmt"
	_ "net/http/pprof"
)

func main() {
	err := convError()
	if err != nil {
		fmt.Println(err)
	}
}

type Foo struct{}

func (f *Foo) Error() string {
	return "error"
}

func convError() error {
	var foo *Foo // foo is nil pointer
	if foo == nil {
		return nil
	}
	return foo
}
