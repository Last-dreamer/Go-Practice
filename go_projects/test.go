package main

import "fmt"

func addTwoNumbers(oneNumber, twoNumber int) int {
	return oneNumber + twoNumber
}

func greeting(name string) string{
	return name
}

func main(){

 

	add := addTwoNumbers(5,95)
	newName := greeting("asim0")

	if newName  == "asim" {
		fmt.Println("yeah it's asim")
	} else {
		fmt.Println("yeah it's not asim")

	} 
	fmt.Println("addTwoNumbers ", add)
	var name = "Single Creation "
	fmt.Println("testing "+ name)
	fmt.Println("hello ", newName)

	var typeInference string  = "it's a type inference .."
	fmt.Println("testing ", typeInference)

	createAndAssign := "create And Assign"
	fmt.Println("testing , ", createAndAssign)

	a, b := 1, "compound creation"
	fmt.Println("testing , ", a, "testing 2 ", b)

	var  (
		anotherInt = 1
		anotherString string = "some string "
	)

	fmt.Println("testing , ", anotherInt, "testing 2 ", anotherString)

}