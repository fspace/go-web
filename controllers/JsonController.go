package controllers

import (
	"encoding/json"
	"fmt"
	json2 "github.com/fspace/go-web/models/jsonmodel"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type JsonController struct {
}

func (c *JsonController) Index() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		jsonFile, err := os.Open("data/post.json")
		if err != nil {
			fmt.Println("Error opening JSON file:", err)
			return
		}
		defer jsonFile.Close()
		jsonData, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			fmt.Println("Error reading JSON data:", err)
			return
		}
		var post json2.Post
		json.Unmarshal(jsonData, &post)
		fmt.Println(post)
		fmt.Fprintln(writer, post)
	}
}

func (c *JsonController) Decoder() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		jsonFile, err := os.Open("data/post.json")
		if err != nil {
			fmt.Println("Error opening JSON file:", err)
			return
		}
		defer jsonFile.Close()
		decoder := json.NewDecoder(jsonFile)
		for {
			var post json2.Post
			err := decoder.Decode(&post)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error decoding JSON:", err)
				return
			}
			fmt.Println(post)
			fmt.Fprintln(writer, post)
		}
	}
}
