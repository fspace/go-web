package controllers

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/fspace/go-web/models"
	"io/ioutil"
	"net/http"
)

type GobController struct {
}

func (c *GobController) Index() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		post := models.Post{Id: 1, Content: "hello  world!", Author: "sha la"}
		store(post, "runtime/post1")
		var postRead models.Post
		load(&postRead, "runtime/post1")
		fmt.Println(postRead)
	}
}

func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0666)
	if err != nil {
		panic(err)
	}
}
func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}
