package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	for i := 0; i < 10; i++ {
		if i == 5 {
			cancel()
		}
		fn(ctx)
	}
}

func fn(ctx context.Context) {
	select {
	default:
		fmt.Println("RUN")
	case <-ctx.Done():
		fmt.Println("Done, close file")
	}
}
