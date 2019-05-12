package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"

	. "github.com/fspace/go-web/middlewares"
	"github.com/fspace/go-web/models"
	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

// 中间件
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		f(writer, request)
	}
}
func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}
func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

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
	tmpl := template.Must(template.ParseFiles("views/layout.html", "views/contact.html"))
	r.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		//t, err := template.ParseFiles("views/layout.html", "views/contact.html")
		//if err != nil {
		//	log.Fatal(err)
		//}
		//t.ExecuteTemplate(w, "layout", "")
		if r.Method != http.MethodPost {
			tmpl.ExecuteTemplate(w, "layout", nil)
			return
		}

		details := models.ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		_ = details

		log.Println(tmpl.Name())
		tmpl.ExecuteTemplate(w, "layout", struct{ Success bool }{true})
	})

	r.HandleFunc("/hello/{title}", helloHandler).Methods("GET")
	r.HandleFunc("/foo", logging(foo))
	r.HandleFunc("/bar", logging(bar))
	r.HandleFunc("/hello", Chain(helloHandler, Method("GET"), Logging()))

	// session cookie 相关：
	r.HandleFunc("/secret", secret)
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logout)

	err := http.ListenAndServe(":85", r)
	if err != nil {
		log.Fatal(err)
	}
}
