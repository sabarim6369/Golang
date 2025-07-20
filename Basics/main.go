package main
import "fmt";
func arrayingo(){
	var a=[2]int{1,2};
	for index,value :=range a{
		fmt.Println(index,value)
	}
	a[1]=20;
	for i:=0;i<len(a);i++{
		fmt.Println(a[i])
	}

	 //slices(Dynamic array)
	c:=[]int{1,2,3}  
	for index,value:=range c{
		fmt.Println(index,value)
	}
	c=append(c,3)
	for index,value:=range c{
		fmt.Println(index,value)
	}

	
}
func main(){
	var a int=10;
	fmt.Println(a)
	arrayingo();
	fmt.Println(2=='2')
}