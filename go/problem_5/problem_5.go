/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package problem5

func problem5() int {
	dividend := 20
	for ; !isDivisibleByAll(dividend, 20); dividend += 20 {

	}

	return dividend
}

func isDivisibleByAll(dividend, divisor int) bool {
	for ; divisor > 1; divisor-- {
		if dividend%divisor != 0 {
			return false
		}
	}

	return true
}
