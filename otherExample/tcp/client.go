package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println("Error connecting to server: " + err.Error())
		return
	}
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)

	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\n")

		if strings.ToUpper(inputInfo) == "Q" {
			return
		}

		// send to server
		_, err = conn.Write([]byte(inputInfo))

		if err != nil {
			return
		}

		// read from server
		buf := [512]byte{}
		n, err := conn.Read(buf[:])

		if err != nil {
			fmt.Println("Error reading from server: ", err)
			return
		}

		fmt.Println(string(buf[:n]))
	}

}
