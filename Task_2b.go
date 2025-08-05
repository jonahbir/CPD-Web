package main

import (
	"fmt"
	"unicode"
)

// IsPalindrome checks if a string is a palindrome (ignoring punctuation and case)
func IsPalindrome(s string) bool {
	var cleaned []rune

	// Normalize: keep only letters/numbers, lowercase
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			cleaned = append(cleaned, unicode.ToLower(r))
		}
	}

	// Compare front and back
	n := len(cleaned)
	for i := 0; i < n/2; i++ {
		if cleaned[i] != cleaned[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	text := "A man, a plan, a canal: Panama"
	fmt.Println(IsPalindrome(text)) // Output: true
}
