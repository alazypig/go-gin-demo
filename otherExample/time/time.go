package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ticker := time.NewTicker(time.Second)

	i := 0

	go func() {
		for {
			i++
			fmt.Println(<-ticker.C)

			if i == 5 {
				ch <- 1
				close(ch)
				break
			}
		}
	}()

	for range ch {
	}
}
