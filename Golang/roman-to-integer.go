package main

import (
	"fmt"
)

// Could be further improved by checking
// the Roman string number from the other way around
func romanToInt(s string) int {

	var result int
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// While this is a completely unnecessary use-case of closures
	// I wanted to practice them a little :)

	updateResultWithDeduction := func(currentByte, nextByte byte, index int) int {
		result += romanMap[nextByte] - romanMap[currentByte]
		index += 1
		return index
	}

	for i := 0; i < len(s); i++ {
		currentByte := s[i]

		if i != len(s)-1 {
			nextByte := s[i+1]
			if preceedsDeductableByte(currentByte, nextByte) {
				i = updateResultWithDeduction(currentByte, nextByte, i)
			} else {
				result += romanMap[currentByte]
			}
		} else {
			result += romanMap[currentByte]
		}
	}

	return result
}

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
func preceedsDeductableByte(currentByte, nextByte byte) bool {
	if currentByte == 73 && (nextByte == 86 || nextByte == 88) {
		return true
	} else if currentByte == 88 && (nextByte == 76 || nextByte == 67) {
		return true
	} else if currentByte == 67 && (nextByte == 68 || nextByte == 77) {
		return true
	}
	return false
}

func main() {
	fmt.Printf("%v", romanToInt("MCDLXXVI"))
}
