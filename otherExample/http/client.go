package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://127.0.0.1:12345/go")
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)

	buf := make([]byte, 1024)

	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println(err)
			return
		} else {
			res := string(buf[:n])
			fmt.Println(res)
			break
		}
	}
}
