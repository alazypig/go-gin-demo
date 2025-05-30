package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "Connected")
	fmt.Println("method: ", r.Method)
	fmt.Println("url: ", r.URL.Path)
	fmt.Println("header: ", r.Header)
	fmt.Println("body: ", r.Body)

	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/go", myHandler)

	http.ListenAndServe("127.0.0.1:12345", nil)
}
