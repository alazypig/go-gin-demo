package main

import (
	"errors"
	"net"
	"reflect"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn}
}

func (c *Client) CallRPC(name string, fPtr interface{}) {
	container := reflect.ValueOf(fPtr).Elem()

	f := func(req []reflect.Value) []reflect.Value {
		reqTransport := NewTransport(c.conn)
		errorHandler := func(err error) []reflect.Value {
			outArgs := make([]reflect.Value, container.Type().NumOut())
			for i := 0; i < len(outArgs)-1; i++ {
				outArgs[i] = reflect.Zero(container.Type().Out(i))
			}

			outArgs[len(outArgs)-1] = reflect.ValueOf(&err).Elem()

			return outArgs
		}

		inArgs := make([]interface{}, 0, len(req))
		for _, arg := range req {
			inArgs = append(inArgs, arg.Interface())
		}

		reqRpc := RPCData{Name: name, Args: inArgs}
		b, err := Encode(reqRpc)
		if err != nil {
			panic(err)
		}

		err = reqTransport.Send(b)
		if err != nil {
			return errorHandler(err)
		}

		resp, err := reqTransport.Read()
		if err != nil {
			return errorHandler(err)
		}

		respDecode, _ := Decode(resp)
		if respDecode.Err != "" {
			return errorHandler(errors.New(respDecode.Err))
		}

		if len(respDecode.Args) == 0 {
			respDecode.Args = make([]interface{}, container.Type().NumOut())
		}

		numOut := container.Type().NumOut()
		outArgs := make([]reflect.Value, numOut)

		for i := 0; i < numOut; i++ {
			if i != numOut-1 {
				if respDecode.Args[i] == nil {
					outArgs[i] = reflect.Zero(container.Type().Out(i))
				} else {
					outArgs[i] = reflect.ValueOf(respDecode.Args[i])
				}
			} else {
				outArgs[i] = reflect.Zero(container.Type().Out(i))
			}
		}

		return outArgs
	}

	container.Set(reflect.MakeFunc(container.Type(), f))
}
