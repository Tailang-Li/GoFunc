package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	closeCh := make(chan struct{})

	go handle(ctx, closeCh)

	select {
	case <-ctx.Done():
		fmt.Printf("result of ctx: %+v\n", ctx.Err().Error())
	}

	<-closeCh
}

func handle(ctx context.Context, closeCh chan struct{}) {

LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("percive ctx done from caller, kill myself, err: %+v\n", ctx.Err().Error())
			break LOOP
		default:
		}
	}

	close(closeCh)
}
