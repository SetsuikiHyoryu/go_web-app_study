package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:9491",
	}

	// handlerForm := func(writer http.ResponseWriter, request *http.Request) {
	// 	request.ParseForm()

	// 	// fmt.Fprintln(writer, request.Form) // map[string][]string

	// 	// 当表单和 url 中含有相同的键名时，PostForm 可只选取表单中的值
	// 	fmt.Fprintln(writer, request.PostForm) // map[string][]string
	// }

	// handlerMultipartForm := func(writer http.ResponseWriter, request *http.Request) {
	// 	request.ParseMultipartForm(1024)
	// 	fmt.Fprintln(writer, request.MultipartForm) // map[string][]string
	// }

	// handlerValue := func(writer http.ResponseWriter, request *http.Request) {
	// 	fmt.Fprintln(writer, request.FormValue("name"))
	// 	fmt.Fprintln(writer, request.PostFormValue("name"))
	// }

	// handlerFile := func(writer http.ResponseWriter, request *http.Request) {
	// 	// 传入的参数为字节数
	// 	request.ParseMultipartForm(1024)
	// 	// 因为文件可能是多个，所以用 0 取出第一个。返回的是指向一个 file header 的指针
	// 	fileHeader := request.MultipartForm.File["filedata"][0]

	// 	// 使用 Open 可以打开文件
	// 	file, _error := fileHeader.Open()

	// 	if _error == nil {
	// 		// ioutil.ReadAll 可以把文件读取为 byte 切片
	// 		data, _error := ioutil.ReadAll(file)

	// 		if _error == nil {
	// 			// 把文件内容写入 response
	// 			fmt.Fprintln(writer, string(data))
	// 		}
	// 	}
	// }

	handlerFormFile := func(writer http.ResponseWriter, request *http.Request) {
		// 更简单地获取第一个文件
		file, _, _error := request.FormFile("filedata")

		if _error == nil {
			// ioutil.ReadAll 可以把文件读取为 byte 切片
			data, _error := ioutil.ReadAll(file)

			if _error == nil {
				// 把文件内容写入 response
				fmt.Fprintln(writer, string(data))
			}
		}
	}

	http.HandleFunc("/api", handlerFormFile)
	server.ListenAndServe()
}
