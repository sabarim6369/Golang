package main
import "fmt"
func sum(a []int){
	for index,value:=range a{
		fmt.Println(index,value)
	}
}
//so it is like a reusable funciton.We can pass single value,multiple and also array toooo.
func sumofgivenval(a... int){
	sum:=0;
	for _,n:=range a{
		sum+=n;
	}
	fmt.Print(sum)

}
func main(){
	a:=[]int{1,2,3,4};
	b:=[]int{5,6,7,8};
	a=append(a,b...);       //used to spread the array as separate values
	fmt.Print(a)
	sum(a)
	sumofgivenval(2)
	sumofgivenval(2,3)
	sumofgivenval(a...)

}