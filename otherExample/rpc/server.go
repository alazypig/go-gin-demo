package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

type Rect struct{}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height

	return nil
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = 2 * (p.Width + p.Height)

	return nil
}

func main() {
	rect := new(Rect)

	rpc.Register(rect) // register the service with the RPC server
	rpc.HandleHTTP()   // set up an HTTP handler for the RPC server

	err := http.ListenAndServe(":12345", nil)

	if err != nil {
		log.Panicln(err)
	}
}
