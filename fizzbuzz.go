package fizzbuzz

import "strconv"

var FizzBuzz = fizzbuzz

func fizzbuzz(n int) string {
	if (n % 15) == 0 {
		return "FizzBuzz"
	}
	if (n % 5) == 0 {
		return "Buzz"
	}
	if (n % 3) == 0 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}
