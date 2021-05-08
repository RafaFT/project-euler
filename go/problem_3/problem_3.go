/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143?
*/

package problem3

import (
	"github.com/rafaft/project-euler/utils"
)

func problem() int {
	dividend := 600851475143
	divisor := utils.NextPrime(-1)

	for dividend != 1 {
		if dividend%divisor == 0 {
			dividend /= divisor
		} else {
			divisor = utils.NextPrime(divisor)
		}
	}

	return divisor
}
