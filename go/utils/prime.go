package utils

import (
	"fmt"
	"math"
)

var (
	ErrNegativeInteger = fmt.Errorf("negative numbers are not supported")
)

func IsPrime(n int) (bool, error) {
	if n < 0 {
		return false, ErrNegativeInteger
	}

	if n <= 1 {
		return false, nil
	}

	if n == 2 {
		return true, nil
	}

	if n%2 == 0 {
		return false, nil
	}

	max := int(math.Sqrt(float64(n)))
	for divisor := 3; divisor <= max; divisor += 2 {
		if n%divisor == 0 {
			return false, nil
		}
	}

	return true, nil
}

func NextPrime(n int) int {
	if n < 2 {
		return 2
	}

	n += 1
	if n%2 == 0 {
		n += 1
	}

	for {
		if isPrime, _ := IsPrime(n); isPrime {
			return n
		}

		n += 2
	}
}
