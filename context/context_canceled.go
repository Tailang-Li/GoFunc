package main

import (
	"context"
	"fmt"
	"time"
)

var (
	testTimeDuration = 1 * time.Second
)

func main() {

	for _, caller := range []func(chan struct{}){
		callerWithContextBackground,
		callerWithContextCanceled,
		callerWithContextTimeout,
	} {
		closeCh := make(chan struct{})
		caller(closeCh)
		<-closeCh
	}
}

func callerWithContextBackground(closeCh chan struct{}) {
	go handler(context.Background(), closeCh)
}

// context canceled and context timeout is the same in this testing because of the calling of cancel()
func callerWithContextCanceled(closeCh chan struct{}) {

	ctx, cancel := context.WithTimeout(context.Background(), testTimeDuration)
	defer cancel()

	go handler(ctx, closeCh)
}

func callerWithContextTimeout(closeCh chan struct{}) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go handler(ctx, closeCh)
}

func handler(ctx context.Context, closeCh chan struct{}) {

	select {
	case <-time.After(testTimeDuration):
		fmt.Println("handler ended with timer time out")
	case <-ctx.Done():
		fmt.Printf("handler perceive context done from caller, reason is: %+v\n", ctx.Err().Error())
	}

	close(closeCh)
}
