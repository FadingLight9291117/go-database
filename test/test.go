package main

import (
	"fmt"
	"strings"
)

func main() {

	stre := "insert 1 as  as "

	words_ := strings.Split(stre, " ")

	// filter
	words := make([]string, 0, len(words_))
	for _, word := range words_ {
		if strings.TrimSpace(word) != "" {
			words = append(words, word)
		}
	}

	for i, v := range words {
		fmt.Printf("%d => %v\n", i, v)
	}
}
