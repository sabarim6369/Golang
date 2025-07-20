package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func Crud() {
	// GET Request
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	check(err)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("GET Response:\n", string(body))
	resp.Body.Close()

	// POST Request
	postData := []byte(`{"title":"foo","body":"bar","userId":1}`)
	resp, err = http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(postData))
	check(err)
	body, _ = io.ReadAll(resp.Body)
	fmt.Println("\nPOST Response:\n", string(body))
	resp.Body.Close()

	// PUT Request
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer(postData))
	check(err)
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	check(err)
	body, _ = io.ReadAll(resp.Body)
	fmt.Println("\nPUT Response:\n", string(body))
	resp.Body.Close()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
