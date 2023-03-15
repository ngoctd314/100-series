package main

import (
	"fmt"
	"reflect"
	"strings"
)

var done bool
var a string

func main() {
	s := []string{"Ngoc", "TD", "TDN"}

	t := reflect.TypeOf(s)
	fmt.Println(t.Size())

	rs := []byte{}
	for _, v := range s {
		fmt.Println(len(v))
		rs = append(rs, v...)
	}
	fmt.Println(len(rs))

	// fmt.Println(concat([]string{"Ngoc", "TD"}))
	// fmt.Println(concatReAllocate([]string{"Ngoc", "TD"}))
	// fmt.Println(concatCustom([]string{"Ngoc", "TD"}))
}

func concat(values []string) string {
	sb := strings.Builder{}
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}

func concatReAllocate(values []string) string {
	rs := ""

	for _, value := range values {
		rs += value
	}

	return rs
}

func concatCustom(values []string) string {
	rs := make([]byte, 0)

	for _, value := range values {
		rs = append(rs, value...)
	}

	return string(rs)
}

func concatCustomUpgrade1(values []string) string {
	n := len(values)
	cnt := 0
	for i := 0; i < n; i++ {
		cnt += len(values[i])
	}

	rs := make([]byte, 0, cnt)

	for _, value := range values {
		rs = append(rs, value...)
	}

	return string(rs)
}

func concatCustomUpgrade2(values []string) string {
	n := len(values)
	cnt := 0
	for i := 0; i < n; i++ {
		cnt += len(values[i])
	}

	sb := strings.Builder{}
	sb.Grow(cnt)

	for _, value := range values {
		sb.WriteString(value)
	}

	return sb.String()
}
