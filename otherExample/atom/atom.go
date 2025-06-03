package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int64
var l sync.Mutex
var wg sync.WaitGroup

func add() {
	x++
	wg.Done()
}

func mutexAdd() {
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

func atomAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		// go add()
		// go mutexAdd()
		go atomAdd()
	}

	wg.Wait()

	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}
