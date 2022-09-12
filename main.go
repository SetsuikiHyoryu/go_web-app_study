package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:9491",
	}

	http.HandleFunc("/write", func(write http.ResponseWriter, response *http.Request) {
		_string := `
			<html>
			  <head><title>Go Web</title></head>
			  <body><h1>Go Web</h1></body>
			</html>
		`

		// Write 被调用时 header 中未设置 Content-Type 的话，数据的前 512 字节会被用来检测 Contetn-Type
		// 如果没有显示调用 WriteHeader，那么会隐式地调用 WriteHeader(http.StatusOK)
		write.Write([]byte(_string))
	})

	http.HandleFunc("/write-header", func(write http.ResponseWriter, response *http.Request) {
		write.WriteHeader(501)
		fmt.Fprintln(write, "No such service, try next door")
	})

	http.HandleFunc("/header", func(write http.ResponseWriter, request *http.Request) {
		write.Header().Set("Location", "http://www.google.co.jp")
		// 注意调用 WriteHeader 后无法再修改 header
		write.WriteHeader(302)
	})

	http.HandleFunc("/json", func(write http.ResponseWriter, request *http.Request) {
		write.Header().Set("Content-Type", "application/json")

		type Post struct {
			User    string
			Threads []string
		}

		post := &Post{
			User:    "hyoryu",
			Threads: []string{"first", "seconde", "third"},
		}

		// json.Marshal 将数据转化为 json 编码
		json, _ := json.Marshal(post)
		write.Write(json)
	})

	server.ListenAndServe()
}
