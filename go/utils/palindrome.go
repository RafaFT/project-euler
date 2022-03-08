package utils

func IsIntPalindrome(n int) bool {
	if n >= 0 && n < 10 {
		return true
	}

	if n < 0 {
		n *= -1
	}

	digits := make([]int, 0, 20)
	for n != 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	for i := 0; i < len(digits)/2; i++ {
		if digits[i] != digits[len(digits)-i-1] {
			return false
		}
	}

	return true
}
