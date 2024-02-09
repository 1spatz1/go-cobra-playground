package stringer

import (
	"math"
	"strconv"
)

func Reverse(input string) (result string) {
	for _, c := range input {
		result = string(c) + result
	}
	return result
}

func Inspect(input string, digits bool) (count int, kind string) {
	if !digits {
		return len(input), "char"
	}
	return inspectNumbers(input), "digit"
}

func inspectNumbers(input string) (count int) {
	for _, c := range input {
		_, err := strconv.Atoi(string(c))
		if err == nil {
			count++
		}
	}
	return count
}

func IsPrime(input string) (result bool) {
	// Parse input to an integer
	number, err := strconv.Atoi(input)

	// return if parsing failed
	if err != nil {
		return false
	}

	// Check if the number is less than 2
	if number < 2 {
		return false
	}

	// Check for factors from 2 to the square root of the number
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false // Number is divisible by i, so it's not prime
		}
	}

	return true // Number is prime if no factors are found
}
