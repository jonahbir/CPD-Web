package main

import (
	"fmt"
	"strings"
	"unicode"
)

// WordFrequencyCount takes a string and returns a map of word frequencies
func WordFrequencyCount(text string) map[string]int {
	wordCount := make(map[string]int)

	// Normalize text: lowercase and remove punctuation
	normalized := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			return unicode.ToLower(r)
		}
		return -1 // drop punctuation
	}, text)

	// Split by whitespace
	words := strings.Fields(normalized)

	// Count each word
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	text := "The quick brown fox jumps over the lazy dog. The fox was quick!"
	result := WordFrequencyCount(text)
	fmt.Println(result)
}
