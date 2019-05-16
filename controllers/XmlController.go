package controllers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type XmlController struct {
}

func (c *XmlController) Index() http.HandlerFunc {
	type Author struct {
		Id   string `xml:"id,attr"`
		Name string `xml:",chardata"`
	}
	type Post struct { //#A
		XMLName xml.Name `xml:"post"`
		Id      string   `xml:"id,attr"`
		Content string   `xml:"content"`
		Author  Author   `xml:"author"`
		Xml     string   `xml:",innerxml"`
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		xmlFile, err := os.Open("data/post.xml")
		if err != nil {
			fmt.Println("Error opening XML file:", err)
			return
		}
		defer xmlFile.Close()
		xmlData, err := ioutil.ReadAll(xmlFile)
		if err != nil {
			fmt.Println("Error reading XML data:", err)
			return
		}
		var post Post
		xml.Unmarshal(xmlData, &post)
		fmt.Println(post)
	}
}
