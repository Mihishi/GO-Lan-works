package main

import "fmt"

func main() {

	var length int
	var width int

	print("Enter the length: ")
	fmt.Scan(&length)

	print("Enter the width: ")
	fmt.Scan(&width)

	area := length * width

	println("Area of the Rectangle is", area)
}
