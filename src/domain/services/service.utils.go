package services

import (
	"strconv"
	"strings"
)

const DiscountByPalindrome float32 = 0.5

func isPalindromeInt(id int) bool {
	intAsString := strconv.FormatUint(uint64(id), 10)
	return isPalindromeString(intAsString)

}

func isPalindromeString(text string) bool {
	equalCaseText := strings.ToLower(text)
	chars := strings.Split(equalCaseText, "")
	stringLen := len(chars)
	lastArrayPos := stringLen - 1
	for i := 0; i <= (lastArrayPos)/2; i++ {
		if chars[i] != chars[lastArrayPos-i] {
			return false
		}
	}
	return true
}
