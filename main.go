package fizzbuzz

import "strconv"

func FizzBuzz(n int) string {
	var result string

	if n % 3 == 0 {
		result += "Fizz"
	}
	if  n % 5 == 0 {
		result += "Buzz"
	}

	if result == "" {
		result = strconv.Itoa(n)
	}
	return result
}

