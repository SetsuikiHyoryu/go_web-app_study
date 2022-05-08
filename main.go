package main

import "net/http"

func main() {
	// 当请求到到根路径时执行回调函数
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Write([]byte("Hello World！"))
	})

	// 设置 web 服务器
	http.ListenAndServe("localhost:9491", nil) // DefaultServeMux. 可以简单地理解为路由器
}
