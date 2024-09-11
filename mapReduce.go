package main

// MapRdeuce program in Go to count the frequency of words in a text file

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// Mapper function
func mapper(text string) map[string]int {
	words := strings.Fields(text)
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}

// Reducer function
func reducer(wordCount map[string]int) map[string]int {
	// Sort by number of occurrences
	keys := make([]string, 0, len(wordCount))
	for key := range wordCount {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return wordCount[keys[i]] > wordCount[keys[j]]
	})

	// Return words with highest frequency
	result := make(map[string]int)
	for _, key := range keys {
		result[key] = wordCount[key]
	}
	return result
}

// Main function
func main() {
	input, err := os.ReadFile("input_data.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	text := string(input)
	text = strings.ToLower(text)
	wordCount := mapper(text)
	wordCount = reducer(wordCount)

	for word, count := range wordCount {
		fmt.Println(word, count)
	}
}
