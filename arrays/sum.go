package sum

func Reduce[T any](collection []T, f func(T, T) T, initialValue T) T {
	result := initialValue
	for _, value := range collection {
		result = f(result, value)
	}
	return result
}

func Sum(numbers []int) int {
	add := func(a, c int) int { return a + c }

	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	var output []int
	for _, numbers := range numbersToSum {
		output = append(output, Sum(numbers))
	}
	return output
}

func SumAllTails(numbersToSum ...[]int) []int {
	sumTail := func(a, c []int) []int {
		if len(c) == 0 {
			return append(a, 0)
		} else {
			tail := c[1:]
			return append(a, Sum(tail))
		}
	}
	return Reduce(numbersToSum, sumTail, []int{})
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	var balance float64
	for _, t := range transactions {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}
	return balance
}
