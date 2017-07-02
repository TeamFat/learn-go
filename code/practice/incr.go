package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

//原子计数器
func main() {

	var ops uint64 = 0
	//goroutines
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// To atomically increment the counter we
				// use `AddUint64`, giving it the memory
				// address of our `ops` counter with the
				// `&` syntax.
				atomic.AddUint64(&ops, 1)

				// Allow other goroutines to proceed.
				runtime.Gosched()
			}
		}()
	}

	// Wait a second to allow some ops to accumulate.
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
