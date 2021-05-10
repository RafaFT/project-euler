/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package problem10

import (
	"github.com/rafaft/project-euler/utils"
)

func problem10() int {
	sum := 0
	prime := utils.NextPrime(-1)

	for prime < 2_000_000 {
		sum += prime
		prime = utils.NextPrime(prime)
	}

	return sum
}
