package main

import "fmt"

func main() {

	var tempC,tempF float64

	print("Enter the Temperature in Celsius: ")
	fmt.Scan(&tempC)

	tempF = (tempC * 9/5) + 32 

	fmt.Print("The entered Celsius Temperature in Fahrenheit is"," ", tempF,"°F")

}