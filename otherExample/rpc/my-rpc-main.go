package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"net"
	"os"
	"runtime/trace"
	"time"
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
	if u, ok := userDB[id]; ok {
		return u, nil
	}

	return User{}, errors.New("user not found")
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

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

	u, err := Query(1)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Println(u)
	}

	u, err = Query(2)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Println(u)
	}
}
