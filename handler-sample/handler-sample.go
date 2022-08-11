package main

import "net/http"

func main() {
	_helloHandler := helloHandler{}
	_aboutHandler := aboutHandler{}

	// 创建服务器
	server := http.Server{
		Addr:    "localhost:9491",
		Handler: nil, // DefaultServeMux
	}

	// 向 DefaultServeMux 注册 handler
	http.Handle("/hello", &_helloHandler)
	http.Handle("/about", &_aboutHandler)
	// 回调函数也可以是一个匿名函数
	http.HandleFunc("/welcome", welcome)
	// http.HandlerFunc 不是一个函数而是一个函数类型，这里是类型转换
	http.Handle("/home", http.HandlerFunc(home))

	http.NotFoundHandler()

	server.ListenAndServe()
}

// 自定义 Handler
type helloHandler struct{}

func (handler *helloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World！"))
}

type aboutHandler struct{}

func (handler *aboutHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("About!"))
}

// Hanlder 函数
func welcome(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Welcome!"))
}

func home(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Home!"))
}
