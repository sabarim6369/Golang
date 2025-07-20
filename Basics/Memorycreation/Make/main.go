package main

import (
	"fmt"
	"sort"
)
func main(){
	//Reallocates memeory accordingly
	a:=make([]int,3);
	a[0]=19
	a[1]=19
	a[2]=19
	//a[3]=20                  //this results in index out of bounds error.
	a=append(a,34,56,7,8,9)     //with this help of append this make realocate memeory as new value came.This incrase memory efficiency
	fmt.Print(a)
	sort.Ints(a);
	fmt.Println(a);
   // Sort descending using sort.Slice
    // sort.Slice(a, func(i, j int) bool {
    //     return a[i] > a[j]  // > means descending
    // })
	sort.Slice(a,func(i,j int)bool{
		return a[i]>a[j]
	})
fmt.Print(a)
 fruits := []string{"banana", "apple", "mango", "pear"}
    sort.Strings(fruits)
    fmt.Println(fruits) 

}