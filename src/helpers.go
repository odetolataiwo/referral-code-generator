package src

import "strings"

// removeCommas removes commas from a string
func removeCommas(word string) string {
	return strings.ReplaceAll(word, ",", "")
}

// Helper function to calculate power
func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	return a * pow(a, b-1)
}
