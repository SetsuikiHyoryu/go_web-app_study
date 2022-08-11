package main

import "net/http"

func main() {
	http.ListenAndServe(":9491", http.FileServer(http.Dir("http_file-server_sample-file")))
}
