package utils

import (
	"math"
	"strconv"
)

func IsIntPalindrome(n int) bool {
	if n >= 0 && n < 10 {
		return true
	}

	s := strconv.Itoa(int(math.Abs(float64(n))))
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
