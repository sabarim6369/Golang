package main
import "fmt";
func main(){
	a:=[]int{1,2,3,4}
	fmt.Println(a)
	var index int=2;     //remove element in 2nd index
	a=append(a[:index],a[index+1:]...)     //as we would get two separate slice of array to spead the second with first use ...
	fmt.Print("After removal::",a)
}