package main

import (
	"fmt"
	"net"

	"gilab.com/pragmaticreviews/golang-gin-poc/otherExample/tcp-stick/proto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println("Error connecting to server: " + err.Error())
		return
	}
	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := fmt.Sprintf("Hello, this is message %d", i)

		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("Error encoding message:", err)
			return
		}

		conn.Write(data)
	}
}
