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
	fmt.Printf(" Type of slice array c is %T\n",c)
	for index,value:=range c{
		fmt.Println(index,value)
	}
	c=append(c,3)            // c[0]=3  ->not possible as this is a slice array4
	c=append(c,2,3,4,5,5)
	for index,value:=range c{
		fmt.Println(index,value)
	}
	data:=[]string{"hello","world","bye","world"}
	data=append(data[1:3])     //End is not inclusive always that is 1-3 measn index 1 ,2 alone
	data=append(data[:3])   //0 to 2
	data=append(data[1:])  //1 to end
	fmt.Println(data)

	
}
func sliceconcepts(){
	a:=[]int{1,2,3};
	// fmt.Println(a)
	// b:=a[1:2]a
	// fmt.Println(b)
	// for index,value:=range b{
	// 	fmt.Print(index,value)
	// }
	a=append(a,5,6,7)              //append returns a new array after updating
	 fmt.Println(a)
	 a=append(a[:3]);
	 fmt.Println(a)

}
func main(){
	// var a int=10;
	// fmt.Println(a)
	// arrayingo();
	sliceconcepts();
	// fmt.Println(2=='2')
}