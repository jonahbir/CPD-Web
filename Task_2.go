package main

import (
	"fmt"
	"strings"
	"unicode"
)


func WordFrequencyCount(text string) map[string]int {
	wordCount := make(map[string]int)


	normalized := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, text)

	words := strings.Fields(normalized)


	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	text :=" this is getting interesting!"
	result := WordFrequencyCount(text)
	fmt.Println(result)
}
