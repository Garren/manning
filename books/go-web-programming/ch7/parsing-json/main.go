package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	{
		jsonFile, err := os.Open("post.json")
		if err != nil {
			fmt.Println("Error opening JSON file: ", err)
			return
		}
		defer jsonFile.Close()

		jsonData, err := ioutil.ReadAll(jsonFile)
		//fmt.Println("data:", jsonData)
		if err != nil {
			fmt.Println("Error reading json data", err)
			return
		}

		var post Post
		// use marshalling when you've got data sitting in a string in memory
		err = json.Unmarshal(jsonData, &post)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(post)
	}
	{
		jsonFile, err := os.Open("post.json")
		if err != nil {
			fmt.Println("Error opening JSON file: ", err)
			return
		}
		defer jsonFile.Close()

		// use decoder if your data is coming from an io.Reader or stream
		decoder := json.NewDecoder(jsonFile)
		for {
			var post Post
			err := decoder.Decode(&post)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error decoding JSON:", err)
				return
			}
			fmt.Println(post)
		}
	}
	{
		post := Post{
			Id:      1,
			Content: "Hello, World!",
			Author: Author{
				Id:   2,
				Name: "Sau Sheong",
			},
			Comments: []Comment{
				Comment{
					Id:      3,
					Content: "have a great day!",
					Author:  "Adam",
				},
				Comment{
					Id:      4,
					Content: "how are you today?",
					Author:  "Betty",
				},
			},
		}
		output, err := json.MarshalIndent(&post, "", "\t\t")
		if err != nil {
			log.Fatal(err)
			return
		}
		err = ioutil.WriteFile("post-out.json", output, 0644)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	{
		post := Post{
			Id:      1,
			Content: "Hello, World!",
			Author: Author{
				Id:   2,
				Name: "Sau Sheong",
			},
			Comments: []Comment{
				Comment{
					Id:      3,
					Content: "have a great day!",
					Author:  "Adam",
				},
				Comment{
					Id:      4,
					Content: "how are you today?",
					Author:  "Betty",
				},
			},
		}
		jsonFile, err := os.Create("post-out-encoded.json")
		if err != nil {
			log.Fatal(err)
			return
		}
		encoder := json.NewEncoder(jsonFile)
		err = encoder.Encode(&post)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
