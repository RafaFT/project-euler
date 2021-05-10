/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
we can see that the 6th prime is 13.

What is the 10001st prime number?
*/

package problem7

import "github.com/rafaft/project-euler/utils"

func problem7() int {
	prime := -1

	for i := 0; i < 10_001; i++ {
		prime = utils.NextPrime(prime)
	}

	return prime
}
