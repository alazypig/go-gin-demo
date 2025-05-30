package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"gilab.com/pragmaticreviews/golang-gin-poc/otherExample/tcp-stick/proto"
)

func process(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error reading from connection:", err)
			break
		}

		fmt.Println("Received:", msg)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go process(conn)
	}
}
