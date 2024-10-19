package main

import "fmt"


func longestDiverseString(a int, b int, c int) string {
	// Initialize an empty result string
	s := ""

	// Loop until we cannot add any more characters
	for a > 0 || b > 0 || c > 0 {
		// We need to decide the character to add based on the remaining counts
		maxChar, maxCount := getMaxChar(a, b, c, s)
		if maxCount == 0 {
			break // No more characters can be added without violating the rules
		}

		// Append the selected character and reduce its count
		s += maxChar
		switch maxChar {
		case "a":
			a--
		case "b":
			b--
		case "c":
			c--
		}
	}

	return s
}

// Helper function to find the character with the maximum count that won't result in 3 consecutive same characters
func getMaxChar(a, b, c int, s string) (string, int) {
	// Get the last two characters of the current string
	lastTwo := ""
	if len(s) >= 2 {
		lastTwo = s[len(s)-2:]
	}

	// We try to pick the character with the highest count that won't violate the rule
	if (a >= b && a >= c && lastTwo != "aa") || (lastTwo == "bb" || lastTwo == "cc") && a> 0 {
		return "a", a
	}
	if (b >= a && b >= c && lastTwo != "bb") || (lastTwo == "aa" || lastTwo == "cc") && b> 0{
		return "b", b
	}
	if (c >= a && c >= b && lastTwo != "cc") || (lastTwo == "aa" || lastTwo == "bb") && c> 0 {
		return "c", c
	}
	return "", 0 // No valid character
}

func main() {
	fmt.Println(longestDiverseString(1, 1, 7)) // Example output: "aabaa"
}
//0 1 5