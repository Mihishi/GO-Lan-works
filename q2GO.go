package main

import "fmt"

func main() {

	var name string

	println("Enter your name: ")

	fmt.Scan(&name)

	println("Hello", name)
}
