package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var text string
	var letter rune

	fmt.Print("Enter a text string: ")
	fmt.Scanln(&text)

	fmt.Print("Enter a letter to count: ")
	fmt.Scanf("%c", &letter)

	
	charCount := utf8.RuneCountInString(text)
	wordCount := len(strings.Fields(text))
	letterCount := countLetterSimple(text, letter)

	
	fmt.Printf("Number of characters: %d\n", charCount)
	fmt.Printf("Number of words: %d\n", wordCount)
	fmt.Printf("Number of occurrences of '%c': %d\n", letter, letterCount)
}



func countLetterSimple(text string, letter rune) int {
	count := 0
	for i := 0; i < len(text); i++ {
		if rune(text[i]) == letter {
			count++
		}
	}
	return count
}
