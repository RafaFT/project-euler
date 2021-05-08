/*
A palindromic number reads the same both ways. The largest palindrome made from
the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package problem4

import "github.com/rafaft/project-euler/utils"

func problem4() int {
	result := 0

	for n1 := 999; n1 > 99; n1-- {
		for n2 := 999; n2 > 99; n2-- {
			if mult := n1 * n2; mult > result && utils.IsIntPalindrome(mult) {
				result = mult
			}
		}
	}

	return result
}
