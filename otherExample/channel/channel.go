package main

import (
	"fmt"
)

func recv(r chan int, w chan int) {
	v := <-r

	fmt.Println("value: ", v)

	w <- 30

	close(w)
}

func main() {
	main_ch := make(chan int, 1)
	recv_ch := make(chan int, 1)

	go recv(recv_ch, main_ch)

	recv_ch <- 10

	for {
		if v, ok := <-main_ch; ok {
			fmt.Println("receive value: ", v)

		} else {
			close(recv_ch)
			break
		}
	}
}
