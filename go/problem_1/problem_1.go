/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package problem1

func problem() int {
	var sum int

	for dividend := 0; dividend < 1000; dividend++ {
		if dividend%3 == 0 || dividend%5 == 0 {
			sum += dividend
		}
	}

	return sum
}
