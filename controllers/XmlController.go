package controllers

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

type Post struct { //#A
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"`

	Comments []Comment `xml:"comments>comment"`
}

type XmlController struct {
}

func (c *XmlController) Index() http.HandlerFunc {

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

// Decode 使用decoder 流式解码
func (c *XmlController) Decode() http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {
		xmlFile, err := os.Open("data/post.xml")
		if err != nil {
			fmt.Println("Error opening XML file:", err)
			return
		}
		defer xmlFile.Close()

		decoder := xml.NewDecoder(xmlFile)
		for {
			t, err := decoder.Token()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error decoding xml to tokens：", err)
				return
			}

			switch se := t.(type) {
			case xml.StartElement:
				if se.Name.Local == "comment" {
					var comment Comment
					decoder.DecodeElement(&comment, &se)
					fmt.Println("the comment is :", comment)
				}
			}
		}

	}
}

func (c *XmlController) Marshal() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		post := Post{
			Id:      "1",
			Content: " Hello World!",
			Author: Author{
				Id:   "2",
				Name: "Sau Sheong",
			},
		}
		//output, err := xml.Marshal(&post)
		output, err := xml.MarshalIndent(&post, "", "\t")
		if err != nil {
			fmt.Println("Error marshalling to XML:", err)
			return
		}
		// writer.Write( output )
		writer.Write([]byte(xml.Header + string(output)))
		/*
			err = ioutil.WriteFile("post.xml", output, 0644)
			if err != nil {
				fmt.Println("Error writing XML to file:", err)
				return
			}
		*/
	}
}
