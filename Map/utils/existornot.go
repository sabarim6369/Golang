package utils
import "fmt";
func Checkexistornot(){
	mp:=make(map[int]string);
	value,exists:=mp[1]
	if exists{
		fmt.Print(value)
	}else{
		fmt.Println("key not found")
	}
}