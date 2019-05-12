package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var r = mux.NewRouter()

var helloHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello !"))
}

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	// @see http://golang.org/pkg/net/http/#ServeMux和http://godoc.org/github.com/gorilla/mux
	// 有差异
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	//tmpl , _ := template.ParseFiles("views/layout.html")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Welcome to my website!")
		//vars := mux.Vars(r)
		//log.Println(vars)

		//data := TodoPageData{
		//	PageTitle: "My TODO list",
		//	Todos: []Todo{
		//		{Title: "Task 1", Done: false},
		//		{Title: "Task 2", Done: true},
		//		{Title: "Task 3", Done: true},
		//	},
		//}
		//tmpl.Execute(w, data)

		// https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0
		// @see https://blog.rubylearning.com/go-web-programming-nesting-templates-f008418c6cc8
		// @see https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/07.4.html#nested-templates
		// @see http://www.josephspurrier.com/how-to-use-template-blocks-in-go-1-6/
		// 从 template 动作 切换到 block   防止未定义的模板 导致不显示（早期难道是err）  尽管可以在base模板中定义默认模板 但感觉还是不如block优雅
		t, err := template.ParseFiles("views/layout.html", "views/index.html") // NOTE 顺序很重要哦
		if err != nil {
			log.Fatal(err)
		}
		t.ExecuteTemplate(w, "layout", "")
	})
	r.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("views/layout.html", "views/contact.html")
		if err != nil {
			log.Fatal(err)
		}
		t.ExecuteTemplate(w, "layout", "")
	})

	r.HandleFunc("/hello/{title}", helloHandler).Methods("GET")

	err := http.ListenAndServe(":85", r)
	if err != nil {
		log.Fatal(err)
	}
}
