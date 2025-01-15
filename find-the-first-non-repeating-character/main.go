package main

import "fmt"

func firstNonRepeatingCharacter(s string) rune {
	charCount := make(map[rune]int)

	// Hitung jumlah kemunculan setiap karakter
	for _, char := range s {
		charCount[char]++
	}

	// Temukan karakter pertama yang tidak berulang
	for _, char := range s {
		if charCount[char] == 1 {
			return char
		}
	}

	return '_' // Kembalikan '_' jika tidak ada karakter non-repeating
}

func main() {
	input := "swiss"
	fmt.Printf("Input: %s\n", input)
	fmt.Printf("First Non-Repeating Character: %c\n", firstNonRepeatingCharacter(input))
}
