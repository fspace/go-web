package controllers

import (
	"encoding/csv"
	"fmt"
	. "github.com/fspace/go-web/models"
	"net/http"
	"os"
	"strconv"
)

type CsvController struct {
}

func (c *CsvController) Index() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fileName := "runtime/posts.csv"
		csvFile, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer csvFile.Close()

		allPosts := []Post{
			Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
			Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
			Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
			Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
		}

		w2 := csv.NewWriter(csvFile)
		for _, post := range allPosts {
			line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
			err := w2.Write(line)
			if err != nil {
				panic(err)
			}
		}
		w2.Flush()

		// 读取
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		r2 := csv.NewReader(file)
		r2.FieldsPerRecord = -1
		records, err := r2.ReadAll()
		if err != nil {
			panic(err)
		}
		fmt.Println(records)

		var posts []Post
		for _, item := range records {
			id, _ := strconv.ParseInt(item[0], 0, 0)
			post := Post{Id: int(id), Content: item[1], Author: item[2]}
			posts = append(posts, post)
		}
		if len(posts) > 0 {
			fmt.Println(posts[0].Id)
			fmt.Println(posts[0].Content)
			fmt.Println(posts[0].Author)
		} else {
			fmt.Println("no posts!")
		}

	}
}
