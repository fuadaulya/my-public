package main

import (
	"fmt"
	"strconv"
)

func compressString(s string) string {
	if len(s) == 0 {
		return ""
	}

	var result string
	count := 1

	for i := 0; i < len(s); i++ {
		// Jika karakter saat ini sama dengan karakter berikutnya, tambahkan count
		if i < len(s)-1 && s[i] == s[i+1] {
			count++
		} else {
			// Tambahkan karakter dan jumlahnya ke result
			result += string(s[i]) + strconv.Itoa(count)
			count = 1 // Reset count
		}
	}

	return result
}

func main() {
	input := "aaabbc"
	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Compressed String: %s\n", compressString(input))
}
