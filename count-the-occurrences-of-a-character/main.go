package main

import "fmt"

func countOccurrences(s string, target rune) int {
	count := 0

	// Hitung jumlah karakter target dalam string
	for _, char := range s {
		if char == target {
			count++
		}
	}

	return count
}

func main() {
	input := "programming"
	target := 'g'
	fmt.Printf("Input: %s, Target: %c\n", input, target)
	fmt.Printf("Occurrences of '%c': %d\n", target, countOccurrences(input, target))
}
