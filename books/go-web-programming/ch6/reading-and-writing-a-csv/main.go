package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "Hello, World!", Author: "Sau Sheong"},
		Post{Id: 2, Content: "Bonjour Monde", Author: "Pierre"},
		Post{Id: 3, Content: "Hola Mundo", Author: "Pedro"},
		Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	// write our posts to disk
	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	// open our file
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read csv records out
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// map to post structs
	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
