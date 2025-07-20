
package main
import "fmt"
type firststruct struct{
    age int
    name string
}
func (f firststruct) greet(){
    fmt.Printf("Hi->Name is %v and age is %v\n",f.name,f.age)
}
func (v firststruct) allname(){
    fmt.Println(v.name,v.age)
    
}
func main() {
  var a firststruct;
  a.name="sabari";
  a.age=20;
  a.greet()
  a.allname()
  fmt.Println(a);
  b:=firststruct{20,"jamesbond"}
  fmt.Println(b)
}