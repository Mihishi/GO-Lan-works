package main

import (
	"fmt"
)

func main() {
	var choice int
	var temp float64

	fmt.Println("Temperature Converter")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Print("Enter your choice (1 or 2): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter temperature in Celsius: ")
		fmt.Scan(&temp)
		fmt.Print("Temperature in Fahrenheit: ", celsiusToFahrenheit(temp))
	case 2:
		fmt.Print("Enter temperature in Fahrenheit: ")
		fmt.Scan(&temp)
		fmt.Print("Temperature in Celsius: ", fahrenheitToCelsius(temp))
	default:
		fmt.Println("Invalid choice. Please enter 1 or 2.")
	}
}

func celsiusToFahrenheit(celsiusTemp float64) float64 {
	return celsiusTemp*9/5 + 32
}

func fahrenheitToCelsius(fahrenheitTemp float64) float64 {
	return (fahrenheitTemp - 32) * 5 / 9
}
