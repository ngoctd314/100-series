package main

import (
	"context"
	"time"
)

type detachContext struct {
	ctx context.Context
}

// Deadline implements context.Context
func (detachContext) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

// Done implements context.Context
func (detachContext) Done() <-chan struct{} {
	return nil
}

// Err implements context.Context
func (detachContext) Err() error {
	return nil
}

// Value implements context.Context
func (d detachContext) Value(key any) any {
	return d.ctx.Value(key)
}

func main() {

}
