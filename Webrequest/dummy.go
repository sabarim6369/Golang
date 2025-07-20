package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
type Data struct{
	Userid int    `json:"userId"`
	Postid int    `json:"id"`
	Title string    `json:"title"`
	Body string     `json:"body"`

}
func Dummy() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	errorhandler(err)
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	errorhandler(err)

	// content := string(data)
	var postdata Data;
	err=json.Unmarshal(data,&postdata);    //map json data to struct
	errorhandler(err);
fmt.Printf("%+v\n", postdata)
	// fmt.Println(content)
}

func errorhandler(err error) {
	if err != nil {
		panic(err)
	}
}
