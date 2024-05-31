package main

import (
	"fmt"
)

func main() {
	var password string
	fmt.Print("Enter password to check its strength: ")
	fmt.Scan(&password)

	var hasUppercase, hasLowercase, hasNumber bool
	const minLength = 8

	if len(password) < minLength {
		fmt.Println("Password is too short. It should be at least 8 characters long.")
	} else {
		for i := 0; i < len(password); i++ {
			c := password[i]
			if c >= 'A' && c <= 'Z' {
				hasUppercase = true
			}
			if c >= 'a' && c <= 'z' {
				hasLowercase = true
			}
			if c >= '0' && c <= '9' {
				hasNumber = true
			}
		}

		if !hasUppercase {
			fmt.Println("Password should contain at least one uppercase letter.")
		}

		if !hasLowercase {
			fmt.Println("Password should contain at least one lowercase letter.")
		}

		if !hasNumber {
			fmt.Println("Password should contain at least one number.")
		}

		if hasUppercase && hasLowercase && hasNumber {
			fmt.Println("Password is strong.")
		}
	}
}
