/*
The sum of the squares of the first ten natural numbers is 385

The square of the sum of the first ten natural numbers is 3025

Hence the difference between the sum of the squares of the first ten natural
numbers and the square of the sum is .

Find the difference between the sum of the squares of the first one hundred
natural numbers and the square of the sum.
*/

package problem6

func problem6() int {
	sumOfSquares := 0
	squareOfSum := 0

	for n := 1; n < 101; n++ {
		squareOfSum += n
		sumOfSquares += n * n
	}

	return (squareOfSum * squareOfSum) - sumOfSquares
}
