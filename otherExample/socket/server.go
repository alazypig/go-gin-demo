package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	go h.run()

	router.HandleFunc("/ws", myWs)
	if err := http.ListenAndServe("127.0.0.1:12345", router); err != nil {
		fmt.Println(err)
	}
}
