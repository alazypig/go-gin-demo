package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
)

type RPCServer struct {
	addr  string
	funcs map[string]reflect.Value
}

func NewServer(addr string) *RPCServer {
	return &RPCServer{
		addr:  addr,
		funcs: make(map[string]reflect.Value),
	}
}

func (s *RPCServer) Register(fnName string, fFunc interface{}) {
	if _, ok := s.funcs[fnName]; ok {
		// 已经存在了
		return
	}

	s.funcs[fnName] = reflect.ValueOf(fFunc)
}

func (s *RPCServer) Execute(req RPCData) RPCData {
	f, ok := s.funcs[req.Name]
	if !ok {
		e := fmt.Sprintf("func %s not found", req.Name)
		log.Println(e)
		return RPCData{Name: req.Name, Args: nil, Err: e}
	}

	log.Printf("func %s called\n", req.Name)

	inArgs := make([]reflect.Value, len(req.Args))
	for i := range req.Args {
		inArgs[i] = reflect.ValueOf(req.Args[i])
	}

	out := f.Call(inArgs)
	outArgs := make([]interface{}, len(out)-1)
	for i := 0; i < len(out)-1; i++ {
		outArgs[i] = out[i].Interface()
	}

	var err string
	if _, ok := out[len(out)-1].Interface().(error); ok {
		err = out[len(out)-1].Interface().(error).Error()
	}

	return RPCData{Name: req.Name, Args: outArgs, Err: err}
}

func (s *RPCServer) Run() {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Printf("listen %s error: %v\n", s.addr, err)
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}

		go func() {
			connTransport := NewTransport(conn)
			for {
				req, err := connTransport.Read()
				if err != nil {
					if err != io.EOF {
						log.Printf("read err: %v\n", err)
						return
					}
				}

				decReq, err := Decode(req)
				if err != nil {
					log.Printf("decode err: %v\n", err)
					return
				}

				resp := s.Execute(decReq)
				b, err := Encode(resp)
				if err != nil {
					log.Printf("encode err: %v\n", err)
					return
				}

				err = connTransport.Send(b)
				if err != nil {
					log.Printf("send err: %v\n", err)
				}
			}
		}()
	}
}
