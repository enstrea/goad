package main

import (
	"context"
	"fmt"
	"time"
)

var (
	chanMap map[string]chan interface{}
)

// 一件事开N个goroutine
// 根据规则选择指定goroutine
// 异步咋做？

func StartGo(ctx context.Context, key string, callback func(val interface{})) {

	ch := make(chan interface{})
	chanMap[key] = ch

	go func() {
		for {
			select {
			case obj := <-ch:
				callback(obj)
			case <-ctx.Done():
				return

			}
		}
	}()
}

func Put(key string) {

}

func StartGoAsync(ctx context.Context, key string, callback func(val interface{})) {

}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	go func() {
		select {
		case ch <- 3:
		case <-time.After(time.Second):
			fmt.Println("timeout")
		}
	}()

	time.Sleep(time.Second * 2)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	go func() {
		fmt.Println(<-ch)
	}()

	time.Sleep(time.Second)
	fmt.Println("stop")
}
