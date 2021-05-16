/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

package problem20

import (
	"math/big"
	"strconv"
)

func problem20() int {
	b := big.NewInt(1)
	for n := 100; n > 1; n-- {
		b.Mul(b, big.NewInt(int64(n)))
	}

	result := 0
	for _, char := range b.String() {
		digit, _ := strconv.Atoi(string(char))
		result += digit
	}

	return result
}
