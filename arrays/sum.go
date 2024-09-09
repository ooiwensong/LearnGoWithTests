package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var output []int

	for _, numbers := range numbersToSum {
		output = append(output, Sum(numbers))
	}
	return output
}

func SumAllTails(numbersToSum ...[]int) []int {
	var output []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			output = append(output, 0)
		} else {
			output = append(output, Sum(numbers[1:]))
		}
	}

	return output
}
