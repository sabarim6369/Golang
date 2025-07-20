package main

import (
	"map/helpers"
	"map/utils"
	"fmt"
	"condition/helper"
	)

//go mod init Map

func main() {
	utils.Checkexistornot()
	helpers.Helperfunction()
	helpers.Dummy()
	helper.Helperfunction()
	//make(map[keytype]valuetype)
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["c++"] = "cpp"
	languages["py"] = "python"
	fmt.Println(languages)
	for key, val := range languages {
		fmt.Println(key, val)
	}
	delete(languages, "js")
	fmt.Println(languages)

	//read a value;
	//  If the key doesn't exist, it returns the zero value (like 0 for int, "" for string, etc.)
	val := languages["py"]
	fmt.Println(val)

	// nestedmaps;
	nestedmaps := make(map[string]map[string]string)
	// var nestedmaps map[string]map[string]string;
	nestedmaps["user"] = map[string]string{
		"name": "sabari",
		"age":  "12",
	}

}
