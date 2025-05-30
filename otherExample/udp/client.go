package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 12345,
	})
	if err != nil {
		fmt.Println("Error connecting to UDP server: ", err)
		return
	}
	defer socket.Close()

	// send data
	sendData := []byte("Hello, world!")
	_, err = socket.Write(sendData)
	if err != nil {
		fmt.Println("Error sending data: ", err)
		return
	}

	// receive data from server
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFrom(data)
	if err != nil {
		fmt.Println("Error reading data: ", err)
		return
	}

	fmt.Printf("recv: %v, addr: %v, count: %v\r\n", string(data[:n]), remoteAddr, n)
}
