package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	words := os.Args[1:]
	list := []string{}
	seen := make(map[string]bool)

	for _, word := range words {
		if !seen[word] {
			seen[word] = true
			list = append(list, word)
		}
	}
	counter := make(map[string]int)
	for _, word := range words {
		counter[word]++
	}

	sort.Slice(list, func(i, j int) bool {
		return counter[list[i]] > counter[list[j]]
	})

	type TagCount struct {
		Tag   string
		Count int
	}
	fmt.Printf("%-20s %s\n", "Tag", "Count")
	for i := 0; i < len(list); i++ {
		if counter[list[i]] > 1 {
			fmt.Printf("%-20s %d\n", list[i], counter[list[i]])
		}
	}
}
