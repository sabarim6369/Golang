package main

import (
	"fmt"
	"io"
	"net/http"
)

func Dummy() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	errorhandler(err)
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	errorhandler(err)
	
	content := string(data)
	fmt.Println(content)
}

func errorhandler(err error) {
	if err != nil {
		panic(err)
	}
}
