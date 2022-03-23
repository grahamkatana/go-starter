package main

import "fmt"

func main(){
	fmt.Println("Hello welcome to command trivia!")
	// variables and datatypes
	var name string = "Graham"
	var age int = 26
	var balance float64 = 100.23
	var isBadProgrammer bool = true
	surname:="Katana"

	//name = "New name"
	//age = 33
	//	balance = 22.90
	//isBadProgrammer = false

	//fmt.Println(name,surname,age, balance,isBadProgrammer)

	fmt.Printf("Hello, %v %v welcome. You are %v years old and your balance is %v. Are you a bad programmer?%v",name, surname,age,balance,isBadProgrammer)


}

//running the file: go run basics.go
//building the file : go build basics.go
// .\basics.exe