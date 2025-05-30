package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte

		n, err := reader.Read(buf[:])

		if err != nil {
			fmt.Println("Error reading from connection: ", err)
			break
		}

		recvStr := string(buf[:n])
		fmt.Println("Received: ", recvStr)
		conn.Write([]byte("Hello, got: " + recvStr))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:12345")

	if err != nil {
		fmt.Println("Error starting TCP server: ", err)
		return
	}

	for {
		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err)
			continue
		}

		go process(conn)
	}
}
