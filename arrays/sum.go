package sum

func Reduce[T, U any](collection []T, f func(U, T) U, initialValue U) U {
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

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

// func BalanceFor(transactions []Transaction, name string) float64 {
// 	adjustBalance := func(currentBalance float64, t Transaction) float64 {
// 		if t.From == name {
// 			return currentBalance - t.Sum
// 		}
// 		if t.To == name {
// 			return currentBalance + t.Sum
// 		}
// 		return currentBalance
// 	}
// 	return Reduce(transactions, adjustBalance, 0.0)
// }

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
}
