package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

func main() {
	conn, err := rpc.DialHTTP("tcp", ":12345")
	if err != nil {
		log.Panic(err)
	}

	ret := 0
	err = conn.Call("Rect.Area", Params{100, 200}, &ret)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Area of rectangle is:", ret)

	err = conn.Call("Rect.Perimeter", Params{100, 200}, &ret)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Perimeter of rectangle is:", ret)
}
