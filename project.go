package main
import "fmt"

func main(){
	var name string
	var age uint
	fmt.Println("+---------------------------------------------------------+")
	fmt.Printf("Enter your name\n>>")
	fmt.Scan(&name)
	fmt.Printf("Hello %v! Welcome lets see if you qualify\n",name)
	fmt.Println("+---------------------------------------------------------+")
	fmt.Printf("Enter your age\n>>")
	fmt.Scan(&age)
	fmt.Println("+---------------------------------------------------------+")

	if age>=3{
		fmt.Println("Let's proceed to the next phase!")
	}else{
		fmt.Println("We can't proceed to the next phase!")
		return

	}
	fmt.Println("Phase one passed!")


}