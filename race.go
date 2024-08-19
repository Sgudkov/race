package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var counters = map[int]int{}
	var totalCount int32
	mu := &sync.Mutex{}
	// This loop launches 5 goroutines, each of which will increment the counters and totalCount
	for i := 0; i < 5; i++ {
		go func(counters map[int]int, th int, mu *sync.Mutex) {
			// Lock the mutex to ensure exclusive access to the counters
			mu.Lock()
			defer mu.Unlock()
			for j := 0; j < 5; j++ {
				counters[th*10+j]++
			}
			//Another way to ensure exclusive access to the counters
			//atomic is used inside of the mutex
			atomic.AddInt32(&totalCount, 1)

		}(counters, i, mu)
	}
	time.Sleep(2 * time.Millisecond)
	mu.Lock()
	fmt.Printf("counters: %v\n", counters)
	mu.Unlock()
	fmt.Printf("totalCount: %d\n", totalCount)
}
