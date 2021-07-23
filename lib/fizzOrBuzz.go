package fizzOrBuzz

import "strconv"

/**
 *	Return fizz/buzz/fizzbuzz
 */
func FizzOrBuzz(n, a, b int, str1, str2 string) string {
	if (n % a == 0 && n % b == 0) {
		return str1+str2
	} else if (n % a == 0) {
		return str1
	} else if (n % b == 0) {
		return str2
	} else {
		return strconv.Itoa(n)
	}
}