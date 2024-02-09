package FibonacciFizzBuzz

import "strconv"

func FibsFizzBuzz(n int) []string {

	var calc = []int{0, 1}
	var sum = 0

	// calculate the numbers of the slice
	for round := 0; round < (n - 1); round++ {
		sum = 0
		for _, num := range calc[len(calc)-2:] {
			sum += num
		}
		calc = append(calc, sum)
	}

	// remove first element, this kata starts with index 1
	calc = calc[1:]

	// iterate over calculated numbers and build the resulting string slice
	var result []string

	for _, value := range calc {
		if value%15 == 0 {
			result = append(result, "FizzBuzz")
			continue
		}

		if value%5 == 0 {
			result = append(result, "Buzz")
			continue
		}

		if value%3 == 0 {
			result = append(result, "Fizz")
			continue
		}

		result = append(result, strconv.Itoa(value))
	}

	return result
}
