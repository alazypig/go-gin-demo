package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 12345,
	})

	if err != nil {
		fmt.Println("Error starting UDP server: ", err)
		return
	}
	defer listen.Close()

	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])

		if err != nil {
			fmt.Println("Error reading from UDP: ", err)
			continue
		}

		fmt.Printf("data: %v, addr: %v, count: %v\r\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP([]byte("Hello, got: "+string(data[:n])), addr)

		if err != nil {
			fmt.Println("Error writing to UDP: ", err)
			continue
		}
	}
}
