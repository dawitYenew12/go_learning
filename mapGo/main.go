package main

import (
	"fmt")

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		firstChar := rune(name[0])
		if _, ok := counts[firstChar]; !ok {
			counts[firstChar] = make(map[string]int)
		}
		counts[firstChar][name]++
	}
	return counts
}

func main() {
	names := []string{"Alice", "Bob", "Charlie", "David", "Amy", "Bob", "David", "Bob", "Charlie", "Alice", "David"}
	counts := getNameCounts(names)
	for firstChar, names := range counts {
		for name, count := range names {
			fmt.Printf("count for [%c] [%v]: %d\n", firstChar, name, count)
		}
	}
}