package main

import "fmt"

func first(w chan<- int) { // single channel
	for i := 0; i < 10; i++ {
		w <- i
	}

	close(w)
}

func second(r <-chan int, w chan<- int) {
	for {
		i, ok := <-r

		if !ok {
			break
		}

		w <- i * 2
	}
	close(w)
}

func main() {
	ch1, ch2 := make(chan int, 1), make(chan int, 1)

	go first(ch1)
	go second(ch1, ch2)

	for i := range ch2 {
		fmt.Println(i)
	}
}
