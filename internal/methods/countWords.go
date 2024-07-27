package methods

import (
	"fmt"
	"strings"
)

func main() {
	text := "this is a test this is only a test"
	words := strings.Fields(text)
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	for word, count := range wordCount {
		fmt.Printf("Word: %s, Count: %d\n", word, count)
	}
}
