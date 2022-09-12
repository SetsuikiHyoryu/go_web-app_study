package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:9491",
	}

	// handlerReadBody := func(writer http.ResponseWriter, request *http.Request) {
	// 	length := request.ContentLength
	// 	body := make([]byte, length)
	// 	request.Body.Read(body)
	// 	fmt.Fprintln(writer, string(body))
	// }

	handlerQuery := func(writer http.ResponseWriter, request *http.Request) {
		url := request.URL
		query := url.Query()

		id := query["id"]
		log.Println(id)

		name := query.Get("name")
		log.Println(name)
	}

	http.HandleFunc("/api", handlerQuery)
	server.ListenAndServe()
}
