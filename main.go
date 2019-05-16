package main

import (
	"github.com/fspace/go-web/controllers"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hi go web"))
}

func main() {

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("public/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", index)

	csvCtrl := controllers.CsvController{}
	mux.HandleFunc("/csv", csvCtrl.Index())

	gobCtrl := controllers.GobController{}
	mux.HandleFunc("/gob", gobCtrl.Index())

	xmlCtrl := controllers.XmlController{}
	mux.HandleFunc("/xml", xmlCtrl.Index())

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
