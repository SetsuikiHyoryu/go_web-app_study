# Go Web 应用学习

## 配置环境

1. 在项目根目录下安装 `go.mod` 文件。

   ```powershell
   go mod init github.com/solenovex/web-tutorial
   ```

## 一个最简单的 web 应用

```console
# 目录结構
E:.
│  go.mod
│  main.go
│  README.md
│
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
