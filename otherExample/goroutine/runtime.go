package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func printKey(key string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println(key, " ", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	go printKey("A", &wg)
	go printKey("B", &wg)

	wg.Wait()

	time.Sleep(time.Second)
}
