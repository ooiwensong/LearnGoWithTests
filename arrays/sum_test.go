package sum

import (
	"slices"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	arr1 := []int{1, 2}
	arr2 := []int{0, 9}

	got := SumAll(arr1, arr2)
	want := []int{3, 9}

	// DeepEqual compares if two variables of same type are deeply equal.
	// In the case of arrays/ slices, they are deeply equal if their corresponding
	// elements are deeply equal.
	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("got %v want %v, given %v and %v", got, want, arr1, arr2)
	// }

	// slices.Equal do a simple shallow compare of slices
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v, given %v and %v", got, want, arr1, arr2)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		arr1 := []int{1, 2}
		arr2 := []int{0, 9}

		got := SumAllTails(arr1, arr2)
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(a, c int) int { return a * c }

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(a, c string) string {
			return a + c
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(riya), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(adil), 175)

	// AssertEqual(t, BalanceFor(transactions, "Riya"), 100)
	// AssertEqual(t, BalanceFor(transactions, "Chris"), -75)
	// AssertEqual(t, BalanceFor(transactions, "Adil"), -25)
}

func TestFind(t *testing.T) {
	type Person struct {
		Name string
	}

	t.Run("find the best programmer", func(t *testing.T) {
		people := []Person{
			Person{Name: "Clark Kent"},
			Person{Name: "Bruce Wayne"},
			Person{Name: "Wally West"},
		}

		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Wally")
		})

		AssertTrue(t, found)
		AssertEqual(t, king, Person{Name: "Wally West"})
	})

	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %+v, want true", got)
	}
}
