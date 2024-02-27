package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func GetJsonExternalAPIExample() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Post ID: %d, Title: %s, Body: %s\n", post.Id, post.Title, post.Body)
}
