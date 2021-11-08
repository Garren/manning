package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
	<?xml version="1.0" encoding="utf-8"?>
	<post id="1">
		<content>Hello, World!</content>
		<author id="2">Sau Sheong</author>
	</post>

	$ go run .
	{{ post} 1 Hello, World! Sau Sheong
		<content>Hello, World!</content>
		<author id="2">Sau Sheong</author>
	}
*/

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"`
}

type Post2 struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   string    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  string `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	{
		xmlFile, err := os.Open("post.xml")
		if err != nil {
			fmt.Println("Error opening XML file: ", err)
			return
		}
		defer xmlFile.Close()

		xmlData, err := ioutil.ReadAll(xmlFile)
		if err != nil {
			fmt.Println("Error reading XML data", err)
			return
		}

		var post Post
		xml.Unmarshal(xmlData, &post)
		fmt.Println(post)
	}
	{
		xmlFile, err := os.Open("post2.xml")
		if err != nil {
			fmt.Println("Error opening XML file: ", err)
			return
		}
		defer xmlFile.Close()

		xmlData, err := ioutil.ReadAll(xmlFile)
		if err != nil {
			fmt.Println("Error reading XML data", err)
			return
		}

		var post Post2
		xml.Unmarshal(xmlData, &post)
		fmt.Println(post)
	}
	{
		xmlFile, err := os.Open("post2.xml")
		if err != nil {
			fmt.Println("Error opening XML file: ", err)
			return
		}
		defer xmlFile.Close()

		// Open returns an io.Reader which we need here
		decoder := xml.NewDecoder(xmlFile)
		for {
			t, err := decoder.Token()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error decoding xml into tokens:", err)
				return
			}
			switch se := t.(type) {
			case xml.StartElement:
				if se.Name.Local == "comment" {
					var comment Comment
					decoder.DecodeElement(&comment, &se)
					fmt.Println(comment)
				}
			}
		}
	}
	{
		post := Post{
			Id:      "1",
			Content: "Hello, World!",
			Author: Author{
				Id:   "2",
				Name: "Sau Sheong",
			},
		}
		//output, err := xml.Marshal(&post)
		output, err := xml.MarshalIndent(&post, "", "\t")
		if err != nil {
			fmt.Println("Error mashalling to XML:", err)
			return
		}
		err = ioutil.WriteFile("post3.xml", []byte(xml.Header+string(output)), 0644)
		if err != nil {
			fmt.Println("Error marshalling XML to file:", err)
			return
		}
	}
	{
		post := Post{
			Id:      "1",
			Content: "Hello, World!",
			Author: Author{
				Id:   "2",
				Name: "Sau Sheong",
			},
		}
		xmlFile, err := os.Create("post4.xml")
		if err != nil {
			fmt.Println("Error mashalling to XML:", err)
			return
		}

		encoder := xml.NewEncoder(xmlFile)
		encoder.Indent("", "\t")
		err = encoder.Encode(&post)
		if err != nil {
			fmt.Println("Error encoding XML to file:", err)
			return
		}
	}
}
