package main

import (
	"fmt"
	"io"
	"os"
)
func main(){
	fmt.Println("Hellow world")
	file,err:=os.Create("./myfile.txt");
	if err!=nil{
		panic(err);
	}
	length,err:=io.WriteString(file,"Helclo world");
	if err!=nil{
		panic(err);

	}
	fmt.Print(length,"is the length of the word written in the file");
}