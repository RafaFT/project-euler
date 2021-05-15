/*
2**15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2**1000?
*/

package problem16

import (
	"math/big"
	"strconv"
)

func problem16() int {
	b := big.NewInt(0)
	s := b.Exp(big.NewInt(2), big.NewInt(1000), nil).String()

	result := 0
	for _, char := range s {
		digit, _ := strconv.Atoi(string(char))
		result += digit
	}

	return result
}
