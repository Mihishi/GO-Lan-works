package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var random int = rand.Intn(10)
	//fmt.Println("Random number is", random)

	var guess int

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)


		if guess < random {
			fmt.Println("It is low! Try again.")
			continue
		} else if guess > random {
			fmt.Println("It is high! Try again.")
			continue
		} else {
			fmt.Printf("Congratulations! You guessed the random number",random)
			break
		}
	}
}
