package controllers

import (
	"html/template"
	"net/http"
)

// 如何处理依赖 ？
// - 共享依赖通过结构类型字段注入
// - 处理器级别的依赖通过函数参数注入
type SiteController struct {
	//db     *someDatabase
	//router *someRouter
	//email  EmailSender
}

/*
@see https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831
func (c *SiteController) handleSomething(...) http.HandlerFunc {
	// thing := prepareThing() // The prepareThing is called only once  尽量只读  如果写 请考虑使用mutex 等保护机制
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}
*/

func (c *SiteController) Index() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("views/layout.html", "views/site/index.html"))
	return func(writer http.ResponseWriter, request *http.Request) {
		tmpl.ExecuteTemplate(writer, "layout", nil)
	}
}
