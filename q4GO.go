package main

import "fmt"

func main() {

	var num1, num2, condition int

	print("Enter the num1: ")
	fmt.Scan(&num1)

	print("Enter the num2: ")
	fmt.Scan(&num2)

	println("Press 1 to addition")
	println("Press 2 to substraction")
	println("Press 3 to multiplication")
	println("Press 4 to division")

	print("Enter the condition: ")
	fmt.Scan(&condition)

	if condition == 1 {
		println(num1 + num2)
	} else if condition == 2 {
		println(num1 - num2)
	} else if condition == 3 {
		println(num1 * num2)
	} else if condition == 4 {
		println(num1 / num2)
	} else {
		println("Invalid")
	}

}
