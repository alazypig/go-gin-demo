package main

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

type User struct {
	Name string
	Age  int
}

var userDB = map[int]User{
	1: User{"Alice", 25},
	3: User{"Bob", 30},
	5: User{"Charlie", 35},
}

func QueryUser(id int) (User, error) {
	span := opentracing.GlobalTracer().StartSpan("QueryUser")
	defer span.Finish()

	span.SetTag("user.id", id)

	if u, ok := userDB[id]; ok {
		span.SetTag("result", "success")
		return u, nil
	}

	span.SetTag("result", true)
	span.LogKV("event", "user not found")

	return User{}, errors.New("user not found")
}

func main() {
	cfg := config.Configuration{
		ServiceName: "my-rpc-server",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "localhost:6831",
		},
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Fatal("Can not initialize jaeger tracer: ", err)
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	// server config
	gob.Register(User{})
	addr := "localhost:12345"
	srv := NewServer(addr)

	// start server
	srv.Register("QueryUser", QueryUser)
	go srv.Run()

	time.Sleep(time.Second)

	// client code

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	cli := NewClient(conn)

	var Query func(int) (User, error)
	cli.CallRPC("QueryUser", &Query)

	spanCtx := opentracing.GlobalTracer().StartSpan("Client.QueryUser")
	opentracing.ContextWithSpan(context.Background(), spanCtx)

	u, err := Query(1)
	if err != nil {
		spanCtx.SetTag("error", true)
		fmt.Printf("error: %v\n", err)
	} else {
		spanCtx.SetTag("result", "success")
		fmt.Println(u)
	}
	spanCtx.Finish()

	spanCtx2 := opentracing.GlobalTracer().StartSpan("Client.QueryUser")
	opentracing.ContextWithSpan(context.Background(), spanCtx2)

	u, err = Query(2)
	if err != nil {
		spanCtx2.SetTag("error", true)
		fmt.Printf("error: %v\n", err)
	} else {
		spanCtx2.SetTag("result", "success")
		fmt.Println(u)
	}
	spanCtx2.Finish()
}
