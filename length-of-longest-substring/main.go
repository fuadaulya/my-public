package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// map untuk melacak indeks terakhir karakter
	charIndex := make(map[rune]int)
	maxLength := 0
	windowStart := 0 // Posisi awal dari window

	// Menelusuri setiap karakter di dalam string
	for windowEnd, char := range s {
		// Jika karakter sudah ada dalam map dan posisinya dalam window
		lastIndex, found := charIndex[char]
		if found && lastIndex >= windowStart {
			// Geser windowStart untuk menghindari duplikasi karakter
			windowStart = lastIndex + 1
		}

		// Update indeks karakter dalam map
		charIndex[char] = windowEnd

		// Hitung panjang window saat ini
		currentLength := windowEnd - windowStart + 1
		// Perbarui panjang maksimum jika diperlukan
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func main() {
	input := "abcabcbb"
	fmt.Println("Input:", input)
	fmt.Println("Output:", lengthOfLongestSubstring(input))
}
