package main
import "fmt"

func main(){
	var name string
	var age int
	fmt.Print("Enter your name>>")
	fmt.Scan(&name)
	fmt.Print("Enter your age>>")
	fmt.Scan(&age)
	fmt.Printf("Hello %v you are %v years old",name,age)

}