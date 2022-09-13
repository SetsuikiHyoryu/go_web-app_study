package main

import (
	"log"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

func main() {
	server := http.Server{
		Addr: "localhost:9491",
	}

	http.HandleFunc("/template", func(write http.ResponseWriter, request *http.Request) {
		_template, _ := template.ParseFiles("template.html")
		_template.Execute(write, "Hello World!")
	})

	http.HandleFunc("/parse-files", func(write http.ResponseWriter, request *http.Request) {
		// 方式一
		// _template, _ := template.ParseFiles("template.html")

		// 方式二
		_template := template.New("template")
		_template, _ = _template.ParseFiles("template.html")

		_template.Execute(write, "Hello World!")
	})

	http.HandleFunc("/parse-glob", func(write http.ResponseWriter, request *http.Request) {
		_template, _ := template.ParseGlob("*.html")
		_template.Execute(write, "Hello World!")
	})

	http.HandleFunc("/execute-template", func(write http.ResponseWriter, request *http.Request) {
		templates, _ := template.ParseFiles("template.html", "template02.html")
		templates.ExecuteTemplate(write, "templatel02.html", "Hello World!")
	})

	/** Full Example */
	templates := loadTemplates()

	http.HandleFunc("/", func(write http.ResponseWriter, request *http.Request) {
		// 第一个字符是 /，所以将其切掉
		fileName := request.URL.Path[1:]
		_template := templates.Lookup(fileName)

		if _template == nil {
			write.WriteHeader(http.StatusNotFound)
			return
		}

		_error := _template.Execute(write, "hello world.")

		if _error != nil {
			log.Fatalln(_error.Error())
		}
	})

	// Actions
	http.HandleFunc("/actions", func(write http.ResponseWriter, request *http.Request) {
		_template, _ := template.ParseFiles("actions.html")
		rand.Seed(time.Now().Unix())
		_template.Execute(write, rand.Intn(10) > 5)
	})

	server.ListenAndServe()
}

func loadTemplates() *template.Template {
	result := template.New("templates")
	_template, _error := result.ParseGlob("*.html")
	template.Must(_template, _error)
	return result
}
