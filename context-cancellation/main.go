package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	const N = 3 // number of goroutines
	var wg sync.WaitGroup

	wg.Add(N) // add the number of goroutines

	ctx, cancel := context.WithCancel(context.Background())
	go func(cancel context.CancelFunc) {
		time.Sleep(time.Millisecond * 3000)
		cancel()
		fmt.Println("call cancel function")
	}(cancel)

	for i := 0; i < N; i++ {
		go func(
			wg *sync.WaitGroup,
			ctx context.Context,
			i int,
		) {
			defer func() {
				wg.Done() // tell a routine is done and decrease WaitGroup count.
				fmt.Printf("done in %d-th goroutine\n", i)
			}()
			for {
				select {
				case <-time.After(time.Millisecond * 500):
					fmt.Printf("running in %d-th goroutine\n", i)
				case <-ctx.Done():
					return
				}
			}
		}(&wg, ctx, i)
	}

	wg.Wait() // for N routines which is done.
	fmt.Println("finish")
}
