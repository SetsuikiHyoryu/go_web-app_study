# Go Web 应用学习

## 配置环境

1. 在项目根目录下安装 `go.mod` 文件。

   ```powershell
   go mod init github.com/solenovex/web-tutorial
   ```

## 01 一个最简单的 web 应用

```console
# 目录结构
E:.
    go.mod
    main.go
    README.md
```

1. 创建 `main.go` 文件。

   ```go
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
   
   ```

2. 运行 `main.go` 文件。

   ```console
   go run main.go
   ```

3. 使用浏览器查看 `9491` 端口确认输出了 `Hello World!`

## 02 处理（Handle） Web 请求

- `http.Handle`
- `http.HandleFunc`

go 使用 Handler 接收一个请求，每个请求都会生成对应的 goroutine 去处理它。

### 创建服务器

#### http.ListenAndServer(address, handler)

|参数|说明|
|-|-|
|address|网络地址，如果为空字符串，即为所有网络接口的 80 端口。|
|handler|如果为 `nil` 即是 `DefaultServeMux`。|

※ `DefaultServeMux` 是一个 multiplexer（多路复用器。可以看作是路由器）。

#### http.Server

是一个结构体，使用这个方式创建服务器定制能力更强。

|字段|说明|
|-|-|
|Addr|网络地址，如果为空字符串，即为所有网络接口的 80 端口。|
|Handler|如果为 `nil` 即是 `DefaultServeMux`。|
|ListenAndServe()|`http.ListenAndServe` 调用的即为此方法。|
|ListenAndServeTLS()|可以应对 https 的请求。 http 包亦可调用此方法。|

### Handler

handler 是一个接口，它定义了一个方法 `ServeHTTP()`。

#### ServeHTTP(HTTPResponseWriter, \*Request)

|参数|说明|
|-|-|
|HTTPResponseWriter|用来写响应的。|
|\*Request|指向 Request 结构体的指中，它保存的是请求。|

#### DefaultServeMux

是一个 Multiplexer（多路复用器），也是一个 Handler。

#### http.Handle

```go
func Handle(pattern string, handler Handler) {}

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

#### http.HandleFunc

Handler 函数就是行为与 Handler 接口类似的函数，即接收与 `ServeHTTP` 方法一样的两个参数。

#### http.HandlerFunc

`http.HandlerFunc` 不是一个方法，而是一个能把函数变为 `HandleFunc` 的函数类型。

### Handler 案例

```console
# 目录结构
E:.
    go.mod
    main.go
    README.md
```

```go
package main

import "net/http"

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

func main() {
	_helloHandler := helloHandler{}
	_aboutHandler := aboutHandler{}

	// 创建服务器
	server := http.Server{
		Addr:    "localhost:9491",
		Handler: nil, // DefaultServeMux
	}

	http.Handle("/hello", &_helloHandler)
	http.Handle("/about", &_aboutHandler)
	// 回调函数也可以是一个匿名函数
	http.HandleFunc("/welcome", welcome)
	// http.HandlerFunc 不是一个函数而是一个函数类型，这里是类型转换
	http.Handle("/home", http.HandlerFunc(home))

	server.ListenAndServe()
}

```
