package main

import (
	"fmt"
	"unicode"
)

func IsPalindrome(s string) bool {
	var cleaned []rune

	
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			cleaned = append(cleaned, unicode.ToLower(r))
		}
	}


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
