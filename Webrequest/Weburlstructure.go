package main
import (
	"fmt"
	"net/url"
)
func Weburl(){
	const rwurl="https://localhost:5000/learn?coursename=react";
	result,err := url.Parse(rwurl);
	if err!=nil{
		panic(err);
	}
	fmt.Println(result.Host)
	fmt.Println(result.ForceQuery)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())
	fmt.Println(result.Query())
}