package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:9491",
	}

	handler := func(writer http.ResponseWriter, request *http.Request) {
		length := request.ContentLength
		body := make([]byte, length)
		request.Body.Read(body)
		fmt.Fprintln(writer, string(body))
	}

	http.HandleFunc("/post", handler)
	server.ListenAndServe()
}
