package main

import (
	"fmt"
	"sync"
)

var x int
var wg sync.WaitGroup
var lock sync.Mutex

func add() {

	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}

	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()

	fmt.Println(x)
}
